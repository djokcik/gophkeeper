package registry

import (
	"gophkeeper/client"
	"gophkeeper/client/service"
)

type ClientServiceRegistry interface {
	GetAuthService() service.AuthService
	GetUserService() service.UserService
	GetLoginPasswordService() service.LoginPasswordService
}

type clientServiceRegistry struct {
	loginService         service.AuthService
	userService          service.UserService
	loginPasswordService service.LoginPasswordService
}

func (r clientServiceRegistry) GetAuthService() service.AuthService {
	return r.loginService
}

func (r clientServiceRegistry) GetUserService() service.UserService {
	return r.userService
}

func (r clientServiceRegistry) GetLoginPasswordService() service.LoginPasswordService {
	return r.loginPasswordService
}

func NewClientServiceRegistry(cfg client.Config) ClientServiceRegistry {
	api := service.NewRpcService(cfg)
	user := service.NewUserService()
	login := service.NewAuthService(api)

	return &clientServiceRegistry{
		loginService:         login,
		userService:          user,
		loginPasswordService: service.NewLoginPasswordService(api, user),
	}
}
