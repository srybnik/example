package service

//go:generate mockgen -destination=service_mock.go -source=service.go -package=service

import (
	"context"
	"example/internal/common/models"
	"example/pkg/echo_service"
)

type Repo interface {
	GetUser(userID string) (*models.User, error)
}

type EchoService struct {
	echo_service.UnimplementedEchoServiceServer
	repo Repo
}

func New(repo Repo) *EchoService {
	return &EchoService{repo: repo}
}

func (s *EchoService) Echo(ctx context.Context, msg *echo_service.SimpleMessage) (*echo_service.SimpleMessage, error) {
	msg.Num = 234234
	msg.ResourceId = "Test"
	return msg, nil
}
