package service

import "gophkeeper/client/repo/models"

type LoginService interface {
	Login(username string, password string) (models.GophUser, error)
	Register(username string, password string) (models.GophUser, error)
}

// Ensure service implements interface
var _ LoginService = (*loginService)(nil)

type loginService struct {
}

func (l loginService) Login(username string, password string) (models.GophUser, error) {
	if username == "test" {
		return models.GophUser{}, ErrInvalidPassword
	}

	return models.GophUser{Username: username, Password: password}, nil
}

func (l loginService) Register(username string, password string) (models.GophUser, error) {
	if username == "test" {
		return models.GophUser{}, ErrDuplicateName
	}

	return models.GophUser{Username: username, Password: password}, nil
}

func NewLoginService() LoginService {
	return &loginService{}
}
