package registry

import (
	"context"
	"github.com/djokcik/gophkeeper/client"
	"github.com/djokcik/gophkeeper/client/service"
	"github.com/djokcik/gophkeeper/client/service/recordservice"
	"github.com/djokcik/gophkeeper/client/storage"
	"github.com/djokcik/gophkeeper/pkg/common"
)

type ClientServiceRegistry interface {
	GetCryptoService() common.CryptoService
	GetSSLConfigService() common.SSLConfigService
	GetStorageService() service.ClientStorageService

	GetAPIService() service.ClientRPCService
	GetAuthService() service.ClientAuthService
	GetUserService() service.ClientUserService

	GetRecordPersonalDataService() recordservice.RecordPersonalDataService
	GetRecordBankCardService() recordservice.RecordBankCardService
	GetRecordTextDataService() recordservice.RecordTextDataService
	GetRecordBinaryDataService() recordservice.RecordBinaryDataService
}

type clientServiceRegistry struct {
	cryptoService    common.CryptoService
	sslConfigService common.SSLConfigService

	apiService     service.ClientRPCService
	authService    service.ClientAuthService
	userService    service.ClientUserService
	storageService service.ClientStorageService

	recordPersonalDataService recordservice.RecordPersonalDataService
	recordBankCardService     recordservice.RecordBankCardService
	recordTextDataService     recordservice.RecordTextDataService
	recordBinaryDataService   recordservice.RecordBinaryDataService
}

func NewClientServiceRegistry(ctx context.Context, cfg client.Config) ClientServiceRegistry {
	crypto := common.NewCryptoService()
	sslConfig := common.NewSSLConfigService()
	localStorage := storage.NewClientLocalStorage(ctx)

	user := service.NewUserService()
	api := service.NewRPCService(cfg, crypto, sslConfig, localStorage)
	clientStorage := service.NewClientStorageService(api, localStorage, user)
	auth := service.NewClientAuthService(api, crypto)
	record := recordservice.NewClientRecordService(crypto)

	return &clientServiceRegistry{
		apiService: api,

		cryptoService:    crypto,
		sslConfigService: sslConfig,
		storageService:   clientStorage,

		authService: auth,
		userService: user,

		recordPersonalDataService: recordservice.NewRecordPersonalDataService(api, user, record),
		recordBankCardService:     recordservice.NewBankCardService(api, user, record),
		recordTextDataService:     recordservice.NewTextDataService(api, user, record),
		recordBinaryDataService:   recordservice.NewBinaryDataService(api, user, record),
	}
}

func (r clientServiceRegistry) GetRecordBankCardService() recordservice.RecordBankCardService {
	return r.recordBankCardService
}

func (r clientServiceRegistry) GetAPIService() service.ClientRPCService {
	return r.apiService
}

func (r clientServiceRegistry) GetStorageService() service.ClientStorageService {
	return r.storageService
}

func (r clientServiceRegistry) GetCryptoService() common.CryptoService {
	return r.cryptoService
}

func (r clientServiceRegistry) GetSSLConfigService() common.SSLConfigService {
	return r.sslConfigService
}

func (r clientServiceRegistry) GetAuthService() service.ClientAuthService {
	return r.authService
}

func (r clientServiceRegistry) GetUserService() service.ClientUserService {
	return r.userService
}

func (r clientServiceRegistry) GetRecordPersonalDataService() recordservice.RecordPersonalDataService {
	return r.recordPersonalDataService
}

func (r clientServiceRegistry) GetRecordTextDataService() recordservice.RecordTextDataService {
	return r.recordTextDataService
}

func (r clientServiceRegistry) GetRecordBinaryDataService() recordservice.RecordBinaryDataService {
	return r.recordBinaryDataService
}
