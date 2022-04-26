package registry

import (
	"gophkeeper/server"
	"gophkeeper/server/service"
	"gophkeeper/server/storage"
)

type ServerServiceRegistry interface {
	GetAuthService() service.AuthService
	GetLoginPasswordService() service.LoginPasswordService
}

type serviceRegistry struct {
	authService          service.AuthService
	loginPasswordService service.LoginPasswordService
}

func NewServerServiceRegistry(cfg server.Config, fileStorage storage.FileStorage) ServerServiceRegistry {
	return &serviceRegistry{
		authService:          service.NewAuthService(cfg, fileStorage),
		loginPasswordService: service.NewLoginPasswordService(cfg, fileStorage),
	}
}

func (r serviceRegistry) GetAuthService() service.AuthService {
	return r.authService
}

func (r serviceRegistry) GetLoginPasswordService() service.LoginPasswordService {
	return r.loginPasswordService
}
