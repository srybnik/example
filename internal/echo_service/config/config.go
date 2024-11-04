package config

type Config struct {
	HttpPort    string
	GrpcPort    string
	MetricsPort string
}

func ParseConfig() (*Config, error) {

	// TODO: cofig in vault

	return &Config{
		HttpPort:    ":80",
		GrpcPort:    ":85",
		MetricsPort: ":81",
	}, nil
}

func validateKeys(cfg *Config) error {

	return nil
}
