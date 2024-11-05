package application

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"example/internal/echo_service/config"
	"example/internal/echo_service/repo"
	"example/internal/echo_service/services/service"
	"example/pkg/echo_service"
)

const (
	AppName = "echo-services"
)

type Application struct {
	cfg    *config.Config
	logger *zerolog.Logger
	repo   *repo.Repo
	srv    *service.EchoService

	//...

	errCh chan error
	wg    sync.WaitGroup
	ready bool
}

func New() *Application {
	return &Application{
		errCh: make(chan error),
	}
}

func (a *Application) Start(ctx context.Context) error {
	if err := a.initConfig(); err != nil {
		return fmt.Errorf("can't init config: %w", err)
	}

	if err := a.initLogger(); err != nil {
		return fmt.Errorf("can't init logger: %w", err)
	}

	a.repo = repo.New()
	a.srv = service.New(a.repo)

	if err := a.startGrpcServer(ctx); err != nil {
		return fmt.Errorf("can't start grpc server: %w", err)
	}

	if err := a.startHttpServer(ctx); err != nil {
		return fmt.Errorf("can't start http server: %w", err)
	}

	//...

	a.logger.Info().Msgf("all systems started successfully")

	return nil
}

func (a *Application) startGrpcServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", a.cfg.GrpcPort)
	if err != nil {
		return fmt.Errorf("error listening on port '%s': %w", a.cfg.GrpcPort, err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			grpc_recovery.StreamServerInterceptor(),
		),
	)

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		<-ctx.Done()
		s.GracefulStop()
		lis.Close()
	}()

	echo_service.RegisterEchoServiceServer(s, a.srv)

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		a.logger.Info().Msgf("starting gRPC server on port %s", a.cfg.GrpcPort)
		if err = s.Serve(lis); err != nil {
			a.errCh <- fmt.Errorf("grpc server exited with error: %w", err)
		}
	}()

	return nil
}

func (a *Application) startHttpServer(ctx context.Context) error {
	conn, err := grpc.NewClient(a.cfg.GrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("error conn to gRPC server: %w", err)
	}

	mux := grpc_runtime.NewServeMux()
	err = echo_service.RegisterEchoServiceHandler(ctx, mux, conn)
	if err != nil {
		return fmt.Errorf("error registering echo service: %w", err)
	}

	server := http.Server{
		Addr:    a.cfg.HttpPort,
		Handler: mux,
	}
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		<-ctx.Done()

		sCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(sCtx)
		conn.Close()
	}()

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		a.logger.Info().Msgf("starting http server on port %s", a.cfg.HttpPort)
		if err = server.ListenAndServe(); err != nil {
			a.errCh <- fmt.Errorf("http server exited with error: %w", err)
		}
	}()

	return nil
}

func (a *Application) initConfig() error {
	cfg, err := config.ParseConfig()
	if err != nil {
		return fmt.Errorf("can't parse config: %w", err)
	}
	a.cfg = cfg

	//...

	return nil
}

func (a *Application) initLogger() error {
	logger := zerolog.New(os.Stdout).With().Str("app", AppName).Timestamp().Logger()
	a.logger = &logger

	//...

	return nil
}

func (a *Application) Wait(ctx context.Context, cancel context.CancelFunc) error {
	var appErr error

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for err := range a.errCh {
			cancel()
			a.logger.Error().Err(err).Send()
			appErr = err
		}
	}()

	<-ctx.Done()
	a.wg.Wait()
	close(a.errCh)
	wg.Wait()

	return appErr
}
