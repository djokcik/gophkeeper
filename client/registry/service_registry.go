package registry

import (
	"context"
	"gophkeeper/client"
	"gophkeeper/client/service"
	"gophkeeper/client/service/recordservice"
	"gophkeeper/client/storage"
	"gophkeeper/pkg/common"
)

type ClientServiceRegistry interface {
	GetCryptoService() common.CryptoService
	GetSSLConfigService() common.SSLConfigService
	GetStorageService() service.ClientStorageService

	GetApiService() service.ClientRpcService
	GetAuthService() service.ClientAuthService
	GetUserService() service.ClientUserService
	GetRecordPersonalDataService() recordservice.RecordPersonalDataService
}

type clientServiceRegistry struct {
	cryptoService    common.CryptoService
	sslConfigService common.SSLConfigService

	apiService     service.ClientRpcService
	authService    service.ClientAuthService
	userService    service.ClientUserService
	storageService service.ClientStorageService

	loginPasswordService recordservice.RecordPersonalDataService
}

func (r clientServiceRegistry) GetApiService() service.ClientRpcService {
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
	return r.loginPasswordService
}

func NewClientServiceRegistry(ctx context.Context, cfg client.Config) ClientServiceRegistry {
	crypto := common.NewCryptoService()
	sslConfig := common.NewSSLConfigService()
	localStorage := storage.NewClientLocalStorage(ctx)

	user := service.NewUserService()
	api := service.NewRpcService(cfg, crypto, sslConfig, localStorage)
	clientStorage := service.NewClientStorageService(api, localStorage, user)
	auth := service.NewClientAuthService(api, crypto)

	return &clientServiceRegistry{
		apiService: api,

		cryptoService:    crypto,
		sslConfigService: sslConfig,
		storageService:   clientStorage,

		authService:          auth,
		userService:          user,
		loginPasswordService: recordservice.NewLoginPasswordService(api, user, crypto),
	}
}
