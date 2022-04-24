package main

import (
	"context"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"gophkeeper/server/rpchandler"
	"net"
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

	addr, err := net.ResolveTCPAddr("tcp", cfg.Address)
	if err != nil {
		log.Fatal().Err(err).Msg("error ResolveTCPAddr")
	}
	// слушаем протокол TCP на объявленном адресе
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal().Err(err).Msg("error ListenTCP")
	}

	rpcServer := rpc.NewServer()
	rpcServer.Register(rpchandler.NewRpcHandler(cfg))

	go func() {
		rpcServer.Accept(listener)
	}()

	<-ctx.Done()
	log.Info().Msg("Shutdown Server ...")
}
