package testhelpers

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"
)

func CreateContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func LoadClientCertificate() *tls.Config {
	cert, err := ioutil.ReadFile(E2EConfig.SSLCert)
	if err != nil {
		log.Fatalf("Couldn't load ssl certificate *.crt: %+v. %v", E2EConfig, err.Error())
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	return &tls.Config{RootCAs: certPool}
}
