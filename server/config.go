package server

import (
	"github.com/caarlos0/env/v6"
	"gophkeeper/pkg/logging"
)

type Config struct {
	Address   string `env:"ADDRESS"`
	StorePath string `env:"STORE_FILE"`
}

func NewConfig() Config {
	cfg := Config{
		Address:   "127.0.0.1:8080",
		StorePath: "/tmp",
	}

	cfg.parseEnv()

	return cfg
}

func (cfg *Config) parseEnv() {
	var err error

	err = env.Parse(cfg)
	if err != nil {
		logging.NewLogger().Fatal().Err(err).Msg("error parse environment")
	}
}
