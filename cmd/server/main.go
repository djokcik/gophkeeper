package main

import (
	"context"
	"crypto/tls"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"gophkeeper/server/rpchandler"
	"net/rpc"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	cfg := server.NewConfig()
	log := logging.NewLogger()

	log.Info().Msgf("config: %+v", cfg)

	cert, _ := tls.LoadX509KeyPair("cert/localhost.crt", "cert/localhost.key")
	conn, err := tls.Listen("tcp", cfg.Address, &tls.Config{Certificates: []tls.Certificate{cert}})

	if err != nil {
		log.Fatal().Err(err).Msg("error ResolveTCPAddr")
	}

	rpcServer := rpc.NewServer()
	rpcServer.Register(rpchandler.NewRpcHandler(cfg))

	go func() {
		rpcServer.Accept(conn)
	}()

	<-ctx.Done()
	log.Info().Msg("Shutdown Server ...")
}
