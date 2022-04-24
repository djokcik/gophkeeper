package server

import (
	"github.com/caarlos0/env/v6"
	"gophkeeper/pkg/logging"
)

type Config struct {
	Address   string `env:"ADDRESS"`
	StoreFile string `env:"STORE_FILE"`
}

func NewConfig() Config {
	cfg := Config{
		Address:   "127.0.0.1:8080",
		StoreFile: "/tmp/gophkeeper-db.json",
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
