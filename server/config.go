package server

import (
	"github.com/caarlos0/env/v6"
	"gophkeeper/pkg/logging"
)

type Config struct {
	Address        string `env:"ADDRESS"`
	StorePath      string `env:"STORE_FILE"`
	PasswordPepper string `env:"PASSWORD_PEPPER"`

	JWTSecretKey string `env:"JWT_SECRET_KEY"`
	DBSecretKey  string `env:"DB_SECRET_KEY"`

	SSLCertPath string `env:"SSL_CERT_PATH" envDefault:"cert/localhost.crt"`
	SSLKeyPath  string `env:"SSL_KEY_PATH" envDefault:"cert/localhost.key"`
}

func NewConfig() Config {
	cfg := Config{
		Address:        "localhost:8080",
		StorePath:      "/tmp",
		PasswordPepper: "pepper",

		DBSecretKey:  "DBSecretKey",
		JWTSecretKey: "JWTSecretKey",
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
