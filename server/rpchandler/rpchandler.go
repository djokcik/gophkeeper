// Package rpchandler is a collection of rpc handlers for use login, get and save data
package rpchandler

import (
	"context"
	"github.com/rs/zerolog"
	"gophkeeper/pkg/common"
	"gophkeeper/pkg/logging"
	"gophkeeper/server"
	"gophkeeper/server/registry"
	"gophkeeper/server/service"
	"gophkeeper/server/storage"
)

// RpcHandler struct for all rpc handlers and require DI dependencies
type RpcHandler struct {
	cfg server.Config

	user service.ServerUserService
	auth common.AuthService

	recordPersonalData service.ServerRecordPersonalDataService
}

// NewRpcHandler constructor for RpcHandler
func NewRpcHandler(cfg server.Config, store storage.Storage) *RpcHandler {
	authUtils := common.NewAuthUtilsService()
	auth := common.NewAuthService(store, cfg.JWTSecretKey, authUtils)
	serviceRegistry := registry.NewServerServiceRegistry(cfg, store, authUtils)

	return &RpcHandler{
		cfg: cfg,

		user: serviceRegistry.GetUserService(),
		auth: auth,

		recordPersonalData: serviceRegistry.GetRecordPersonalDataService(),
	}
}

func (h *RpcHandler) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "RpcHandler").Logger()

	return &logger
}
