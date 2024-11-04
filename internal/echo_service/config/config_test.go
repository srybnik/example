package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig_ConfigPathNotSet(t *testing.T) {
	t.Setenv(configPath, "")
	cfg, err := ParseConfig()
	assert.Nil(t, cfg)
	assert.EqualError(t, err, "env not set: CONFIG_PATH")
}
