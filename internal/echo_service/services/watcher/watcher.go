package watcher

//go:generate mockgen -destination=watcher_mock.go -source=watcher.go -package=watcher

import (
	"context"
	"time"

	"github.com/rs/zerolog"
)

type Metrics interface {
	SetS3FreeSpace(space int64)
}

type Storage interface {
	UsedSpace(ctx context.Context) (int64, error)
}

type S3SpaceWatcher struct {
	storage    Storage
	metrics    Metrics
	logger     *zerolog.Logger
	totalSpace int64
}

func NewS3SpaceWatcher(storage Storage, metrics Metrics, logger *zerolog.Logger, totalSpace int64) *S3SpaceWatcher {
	return &S3SpaceWatcher{
		storage:    storage,
		metrics:    metrics,
		logger:     logger,
		totalSpace: totalSpace,
	}
}

func (s *S3SpaceWatcher) Start(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	s.checkS3Space(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.checkS3Space(ctx)
		}
	}
}

func (s *S3SpaceWatcher) checkS3Space(ctx context.Context) {
	usedSpace, err := s.storage.UsedSpace(ctx)
	if err != nil {
		s.logger.Error().Msgf("can't get used space from storage: %v", err)
		return
	}

	freeSpace := s.totalSpace - usedSpace
	s.metrics.SetS3FreeSpace(freeSpace)
}
