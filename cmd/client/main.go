package main

import (
	"context"
	"gophkeeper/client"
	"gophkeeper/client/controller"
	"gophkeeper/client/registry"
	"os/signal"
	"syscall"
)

// TODO: не забыть добавить версию
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
