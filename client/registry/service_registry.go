package registry

import (
	"context"
	"github.com/djokcik/gophkeeper/client"
	"github.com/djokcik/gophkeeper/client/service"
	"github.com/djokcik/gophkeeper/client/service/recordservice"
	"github.com/djokcik/gophkeeper/client/storage"
	"github.com/djokcik/gophkeeper/pkg/common"
)

// ClientServiceRegistry is registered all services
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
	auth := service.NewClientAuthService(api)
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

// GetRecordBankCardService returns RecordBankCardService
func (r clientServiceRegistry) GetRecordBankCardService() recordservice.RecordBankCardService {
	return r.recordBankCardService
}

// GetAPIService returns ClientRPCService
func (r clientServiceRegistry) GetAPIService() service.ClientRPCService {
	return r.apiService
}

// GetStorageService returns ClientStorageService
func (r clientServiceRegistry) GetStorageService() service.ClientStorageService {
	return r.storageService
}

// GetCryptoService returns CryptoService
func (r clientServiceRegistry) GetCryptoService() common.CryptoService {
	return r.cryptoService
}

// GetSSLConfigService returns SSLConfigService
func (r clientServiceRegistry) GetSSLConfigService() common.SSLConfigService {
	return r.sslConfigService
}

// GetAuthService returns ClientAuthService
func (r clientServiceRegistry) GetAuthService() service.ClientAuthService {
	return r.authService
}

// GetUserService returns ClientUserService
func (r clientServiceRegistry) GetUserService() service.ClientUserService {
	return r.userService
}

// GetRecordPersonalDataService returns RecordPersonalDataService
func (r clientServiceRegistry) GetRecordPersonalDataService() recordservice.RecordPersonalDataService {
	return r.recordPersonalDataService
}

// GetRecordTextDataService returns RecordTextDataService
func (r clientServiceRegistry) GetRecordTextDataService() recordservice.RecordTextDataService {
	return r.recordTextDataService
}

// GetRecordBinaryDataService returns RecordBinaryDataService
func (r clientServiceRegistry) GetRecordBinaryDataService() recordservice.RecordBinaryDataService {
	return r.recordBinaryDataService
}
