package registry

import (
	"gophkeeper/pkg/common"
	"gophkeeper/server"
	"gophkeeper/server/service"
	"gophkeeper/server/storage"
)

type ServerServiceRegistry interface {
	GetUserService() service.ServerUserService
	GetKeyLockService() service.KeyLockService

	GetRecordPersonalDataService() service.ServerRecordPersonalDataService
}

type serviceRegistry struct {
	user               service.ServerUserService
	recordPersonalData service.ServerRecordPersonalDataService
	keyLock            service.KeyLockService
}

func NewServerServiceRegistry(cfg server.Config, store storage.Storage, auth common.AuthUtilsService) ServerServiceRegistry {
	keyLockService := service.NewStringKeyLock()

	return &serviceRegistry{
		user:               service.NewAuthService(cfg, store, auth),
		recordPersonalData: service.NewServerRecordPersonalDataService(cfg, store, keyLockService),
		keyLock:            keyLockService,
	}
}

func (r serviceRegistry) GetUserService() service.ServerUserService {
	return r.user
}

func (r serviceRegistry) GetKeyLockService() service.KeyLockService {
	return r.keyLock
}

func (r serviceRegistry) GetRecordPersonalDataService() service.ServerRecordPersonalDataService {
	return r.recordPersonalData
}
