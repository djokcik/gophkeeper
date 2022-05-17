package testhelpers

import (
	"github.com/caarlos0/env/v6"
	"log"
)

var E2EConfig config

type config struct {
	Address   string `env:"ADDRESS,required"`
	SSLCert   string `env:"SSL_CERT_PATH,required"`
	StorePath string `env:"STORE_FILE,required"`
}

func (cfg *config) parseEnv() {
	var err error

	err = env.Parse(cfg)
	if err != nil {
		log.Fatal("error parse environment: ", err.Error())
	}
}

func init() {
	E2EConfig.parseEnv()
}
