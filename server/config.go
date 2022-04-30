package server

import (
	"github.com/caarlos0/env/v6"
	"gophkeeper/pkg/logging"
)

type Config struct {
	Address   string `env:"ADDRESS"`
	StorePath string `env:"STORE_FILE"`

	SSLCert string `env:"SSL_CERT" envDefault:"cert/localhost.crt"`
	SSLKey  string `env:"SSL_KEY" envDefault:"cert/localhost.key"`
}

func NewConfig() Config {
	cfg := Config{
		Address:   "localhost:8080",
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
