// Package rpchandler is a collection of rpc handlers for use login, get and save data
package rpchandler

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"gophkeeper/server/registry"
	"gophkeeper/server/storage"
)

// RpcHandler struct for all rpc handlers and require DI dependencies
type RpcHandler struct {
	cfg             server.Config
	serviceRegistry registry.ServerServiceRegistry
}

// NewRpcHandler constructor for RpcHandler
func NewRpcHandler(cfg server.Config) *RpcHandler {
	fileStorage := storage.NewFileStorage(cfg)

	return &RpcHandler{
		cfg:             cfg,
		serviceRegistry: registry.NewServerServiceRegistry(cfg, fileStorage),
	}
}

func (h *RpcHandler) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "RpcHandler").Logger()

	return &logger
}
