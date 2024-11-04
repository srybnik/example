package sender

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SenderSuite struct {
	suite.Suite
	publisher *MockPublisher
	sender    *Sender
}

func TestSender(t *testing.T) {
	suite.Run(t, &SenderSuite{})
}

func (s *SenderSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.publisher = NewMockPublisher(ctrl)

	sender, err := New(s.publisher)
	s.NoError(err)
	s.sender = sender

}

func (s *SenderSuite) TestSend_Ok() {

	msg := []byte("hello world")
	s.publisher.EXPECT().Publish(gomock.Any(), "topic", msg).Return(nil)

	err := s.sender.Send(context.Background(), msg)
	s.NoError(err)
}
