package common

import (
	"crypto/tls"
	"crypto/x509"
	"gophkeeper/client"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"io/ioutil"
)

type SSLConfigService interface {
	LoadClientCertificate(cfg client.Config) (*tls.Config, error)
	LoadServerCertificate(cfg server.Config) (*tls.Config, error)
}

type sslConfigService struct {
}

func NewSSLConfigService() SSLConfigService {
	return &sslConfigService{}
}

func (s sslConfigService) LoadClientCertificate(cfg client.Config) (*tls.Config, error) {
	log := logging.NewFileLogger()

	cert, err := ioutil.ReadFile(cfg.SSLCert)
	if err != nil {
		log.Error().Err(err).Msgf("Couldn't load ssl certificate *.crt: %v", cfg)
		return nil, err
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	return &tls.Config{RootCAs: certPool}, nil
}

func (s sslConfigService) LoadServerCertificate(cfg server.Config) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(cfg.SSLCertPath, cfg.SSLKeyPath)
	if err != nil {
		logging.NewLogger().Err(err).Msgf("Couldn't load ssl certificate *.crt, *.key: %v", cfg)
		return nil, err
	}

	return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
}
