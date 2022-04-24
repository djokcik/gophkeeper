package registry

import "gophkeeper/client/service"

type ServiceRegistry interface {
	GetLoginService() service.LoginService
	GetUserService() service.UserService
}

type serviceRegistry struct {
	loginService service.LoginService
	userService  service.UserService
}

func (r serviceRegistry) GetLoginService() service.LoginService {
	return r.loginService
}

func (r serviceRegistry) GetUserService() service.UserService {
	return r.userService
}

func NewServiceRegistry() ServiceRegistry {
	return &serviceRegistry{
		loginService: service.NewLoginService(),
		userService:  service.NewUserService(),
	}
}
