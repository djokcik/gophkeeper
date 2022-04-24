package client

import (
	"github.com/caarlos0/env/v6"
	"gophkeeper/pkg/logging"
)

type Config struct {
	Address string `env:"SERVER_ADDRESS" envDefault:"127.0.0.1:8080"`
}

func NewConfig() Config {
	var cfg Config

	cfg.parseEnv()

	return cfg
}

func (cfg *Config) parseEnv() {
	var err error

	err = env.Parse(cfg)
	if err != nil {
		logging.NewLogger().Fatal().Err(err).Msg("ClientConfig: error parse environment")
	}
}
