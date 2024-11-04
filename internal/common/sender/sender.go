package sender

//go:generate mockgen -destination=sender_mock.go -source=sender.go -package=sender

import (
	"context"
	"fmt"
)

type Publisher interface {
	Publish(ctx context.Context, subject string, msg []byte) error
}

type Sender struct {
	publisher Publisher
}

func New(publisher Publisher) (*Sender, error) {
	return &Sender{
		publisher: publisher,
	}, nil
}

func (s *Sender) Send(ctx context.Context, msg []byte) error {

	if err := s.publisher.Publish(ctx, "topic", msg); err != nil {
		return fmt.Errorf("cannot publish notice: %w", err)
	}

	return nil
}
