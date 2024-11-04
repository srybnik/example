package config

import (
	"fmt"
)

type DBCreds struct {
	Host     string `env:"PG_HOST" required:"true"`
	Port     int64  `env:"PG_PORT" required:"true"`
	Username string `env:"PG_USER" required:"true"`
	Password string `env:"PG_PASS" required:"true"`
	Database string `env:"PG_DBNAME" required:"true"`
}

func (cfg DBCreds) BuildDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
	)
}
