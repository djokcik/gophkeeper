package registry

import (
	"github.com/djokcik/gophkeeper/pkg/common"
	"github.com/djokcik/gophkeeper/server"
	"github.com/djokcik/gophkeeper/server/service"
	"github.com/djokcik/gophkeeper/server/storage"
)

// ServerServiceRegistry is registered all server services
type ServerServiceRegistry interface {
	GetUserService() service.ServerUserService
	GetKeyLockService() service.KeyLockService

	GetRecordPersonalDataService() service.ServerRecordPersonalDataService
	GetRecordBankCardService() service.ServerRecordBankCardDataService
	GetRecordTextDataService() service.ServerRecordTextDataService
	GetRecordBinaryDataService() service.ServerRecordBinaryDataService
}

type serviceRegistry struct {
	user    service.ServerUserService
	keyLock service.KeyLockService

	recordPersonalData service.ServerRecordPersonalDataService
	recordBankCard     service.ServerRecordBankCardDataService
	recordTextData     service.ServerRecordTextDataService
	recordBinaryData   service.ServerRecordBinaryDataService
}

func NewServerServiceRegistry(cfg server.Config, store storage.Storage, auth common.AuthUtilsService) ServerServiceRegistry {
	keyLockService := service.NewStringKeyLock()
	recordService := service.NewRecordService(keyLockService, store)

	return &serviceRegistry{
		user:               service.NewAuthService(cfg, store, auth),
		recordPersonalData: service.NewServerRecordPersonalDataService(cfg, recordService),
		recordBankCard:     service.NewServerRecordBankCardDataService(cfg, recordService),
		recordTextData:     service.NewServerRecordTextDataService(cfg, recordService),
		recordBinaryData:   service.NewServerRecordBinaryDataService(cfg, recordService),
		keyLock:            keyLockService,
	}
}

// GetUserService returns ServerUserService
func (r serviceRegistry) GetUserService() service.ServerUserService {
	return r.user
}

// GetKeyLockService returns KeyLockService
func (r serviceRegistry) GetKeyLockService() service.KeyLockService {
	return r.keyLock
}

// GetRecordPersonalDataService returns ServerRecordPersonalDataService
func (r serviceRegistry) GetRecordPersonalDataService() service.ServerRecordPersonalDataService {
	return r.recordPersonalData
}

// GetRecordBankCardService returns ServerRecordBankCardDataService
func (r serviceRegistry) GetRecordBankCardService() service.ServerRecordBankCardDataService {
	return r.recordBankCard
}

// GetRecordTextDataService returns ServerRecordTextDataService
func (r serviceRegistry) GetRecordTextDataService() service.ServerRecordTextDataService {
	return r.recordTextData
}

// GetRecordBinaryDataService returns ServerRecordBinaryDataService
func (r serviceRegistry) GetRecordBinaryDataService() service.ServerRecordBinaryDataService {
	return r.recordBinaryData
}
