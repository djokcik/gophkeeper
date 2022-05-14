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
	GetRecordBankCardService() service.ServerRecordBankCardDataService
	GetRecordTextDataService() service.ServerRecordTextDataService
}

type serviceRegistry struct {
	user    service.ServerUserService
	keyLock service.KeyLockService

	recordPersonalData service.ServerRecordPersonalDataService
	recordBankCard     service.ServerRecordBankCardDataService
	recordTextData     service.ServerRecordTextDataService
}

func NewServerServiceRegistry(cfg server.Config, store storage.Storage, auth common.AuthUtilsService) ServerServiceRegistry {
	keyLockService := service.NewStringKeyLock()

	return &serviceRegistry{
		user:               service.NewAuthService(cfg, store, auth),
		recordPersonalData: service.NewServerRecordPersonalDataService(cfg, store, keyLockService),
		recordBankCard:     service.NewServerRecordBankCardDataService(cfg, store, keyLockService),
		recordTextData:     service.NewServerRecordTextDataService(cfg, store, keyLockService),
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

func (r serviceRegistry) GetRecordBankCardService() service.ServerRecordBankCardDataService {
	return r.recordBankCard
}

func (r serviceRegistry) GetRecordTextDataService() service.ServerRecordTextDataService {
	return r.recordTextData
}
