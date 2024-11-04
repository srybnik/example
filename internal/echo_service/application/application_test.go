package application

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ApplicationSuite struct {
	suite.Suite
	app       *Application
	testError error
}

func TestApplication(t *testing.T) {
	suite.Run(t, &ApplicationSuite{})
}

func (s *ApplicationSuite) SetupTest() {
	s.app = New()
}

func (s *ApplicationSuite) TestStart_ConfigError() {
	ctx := context.Background()
	err := s.app.Start(ctx)
	s.ErrorContains(err, "can't init config:")
}

//TODO: make test
