package registry

import (
	"gophkeeper/client"
	"gophkeeper/client/service"
	"gophkeeper/client/service/recordservice"
	"gophkeeper/pkg/common"
)

type ClientServiceRegistry interface {
	GetCryptoService() common.CryptoService
	GetSSLConfigService() common.SSLConfigService

	GetAuthService() service.AuthService
	GetUserService() service.UserService
	GetRecordPersonalDataService() recordservice.RecordPersonalDataService
}

type clientServiceRegistry struct {
	cryptoService    common.CryptoService
	sslConfigService common.SSLConfigService

	loginService         service.AuthService
	userService          service.UserService
	loginPasswordService recordservice.RecordPersonalDataService
}

func (r clientServiceRegistry) GetCryptoService() common.CryptoService {
	return r.cryptoService
}

func (r clientServiceRegistry) GetSSLConfigService() common.SSLConfigService {
	return r.sslConfigService
}

func (r clientServiceRegistry) GetAuthService() service.AuthService {
	return r.loginService
}

func (r clientServiceRegistry) GetUserService() service.UserService {
	return r.userService
}

func (r clientServiceRegistry) GetRecordPersonalDataService() recordservice.RecordPersonalDataService {
	return r.loginPasswordService
}

func NewClientServiceRegistry(cfg client.Config) ClientServiceRegistry {
	crypto := common.NewCryptoService()
	sslConfig := common.NewSSLConfigService()

	api := service.NewRpcService(cfg, crypto, sslConfig)
	user := service.NewUserService()
	login := service.NewAuthService(api, crypto)

	return &clientServiceRegistry{
		cryptoService:    crypto,
		sslConfigService: sslConfig,

		loginService:         login,
		userService:          user,
		loginPasswordService: recordservice.NewLoginPasswordService(api, user, crypto),
	}
}
