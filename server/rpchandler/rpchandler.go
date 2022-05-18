// Package rpchandler is a collection of rpc handlers for use login, get and save data
package rpchandler

import (
	"context"
	"github.com/djokcik/gophkeeper/pkg/common"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/djokcik/gophkeeper/server"
	"github.com/djokcik/gophkeeper/server/registry"
	"github.com/djokcik/gophkeeper/server/service"
	"github.com/djokcik/gophkeeper/server/storage"
	"github.com/rs/zerolog"
)

// RPCHandler struct for all rpc handlers and require DI dependencies
type RPCHandler struct {
	user service.ServerUserService
	auth common.AuthService

	recordPersonalData service.ServerRecordPersonalDataService
	recordBankCard     service.ServerRecordBankCardDataService
	recordTextData     service.ServerRecordTextDataService
	recordBinaryData   service.ServerRecordBinaryDataService
}

// NewRPCHandler constructor for RPCHandler
func NewRPCHandler(cfg server.Config, store storage.Storage) *RPCHandler {
	authUtils := common.NewAuthUtilsService()
	auth := common.NewAuthService(store, cfg.JWTSecretKey, authUtils)
	serviceRegistry := registry.NewServerServiceRegistry(cfg, store, authUtils)

	return &RPCHandler{
		user: serviceRegistry.GetUserService(),
		auth: auth,

		recordPersonalData: serviceRegistry.GetRecordPersonalDataService(),
		recordBankCard:     serviceRegistry.GetRecordBankCardService(),
		recordTextData:     serviceRegistry.GetRecordTextDataService(),
		recordBinaryData:   serviceRegistry.GetRecordBinaryDataService(),
	}
}

func (h *RPCHandler) Log(ctx context.Context) *zerolog.Logger {
	_, logger := logging.GetCtxLogger(ctx)
	logger = logger.With().Str(logging.ServiceKey, "RPCHandler").Logger()

	return &logger
}
