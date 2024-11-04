package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseCfg(t *testing.T) {
	cfg := DBCreds{
		Host:     "localhost",
		Port:     5432,
		Username: "user",
		Password: "pass",
		Database: "db",
	}

	assert.Equal(t, cfg.BuildDSN(), "host=localhost port=5432 user=user password=pass dbname=db sslmode=disable")
}
