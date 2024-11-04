package watcher

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

type S3SpaceWatcherSuite struct {
	suite.Suite

	metrics *MockMetrics
	storage *MockStorage
	logs    *zerolog.Logger

	s3SpaceWatcher *S3SpaceWatcher

	ctx       context.Context
	testError error
}

func TestBackup(t *testing.T) {
	suite.Run(t, &S3SpaceWatcherSuite{})
}

func (s *S3SpaceWatcherSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.metrics = NewMockMetrics(ctrl)
	s.storage = NewMockStorage(ctrl)

	logs := zerolog.New(os.Stdout).With().Timestamp().Logger()
	s.logs = &logs

	s.s3SpaceWatcher = NewS3SpaceWatcher(s.storage, s.metrics, s.logs, 1000)

	s.ctx = context.Background()
	s.testError = errors.New("test error")
}

func (s *S3SpaceWatcherSuite) TestCheckS3Space_Ok() {
	s.storage.EXPECT().UsedSpace(s.ctx).Return(int64(200), nil)
	s.metrics.EXPECT().SetS3FreeSpace(int64(800))

	s.s3SpaceWatcher.checkS3Space(s.ctx)
}
