package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"example/internal/echo_service/application"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	app := application.New()

	err := app.Start(ctx)
	if err != nil {
		log.Fatalf("Can't start application: %s", err)
	}

	err = app.Wait(ctx, cancel)
	if err != nil {
		log.Fatalf("All systems closed with errors. LastError: %s", err)
	}

	log.Println("All systems closed without errors")
}
