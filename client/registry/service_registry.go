package registry

import (
	"gophkeeper/client"
	"gophkeeper/client/service"
)

type ServiceRegistry interface {
	GetAuthService() service.AuthService
	GetUserService() service.UserService
}

type serviceRegistry struct {
	loginService service.AuthService
	userService  service.UserService
}

func (r serviceRegistry) GetAuthService() service.AuthService {
	return r.loginService
}

func (r serviceRegistry) GetUserService() service.UserService {
	return r.userService
}

func NewServiceRegistry(cfg client.Config) ServiceRegistry {
	api := service.NewRpcService(cfg)

	return &serviceRegistry{
		loginService: service.NewAuthService(api),
		userService:  service.NewUserService(),
	}
}
