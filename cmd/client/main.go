package main

import (
	"context"
	"github.com/djokcik/gophkeeper/client"
	"github.com/djokcik/gophkeeper/client/controller"
	"github.com/djokcik/gophkeeper/client/registry"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	cfg := client.NewConfig()

	serviceRegistry := registry.NewClientServiceRegistry(ctx, cfg)

	go func() {
		controller.Start(ctx, serviceRegistry)
		cancel()
	}()

	<-ctx.Done()
}
