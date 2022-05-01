package service

import (
	"gophkeeper/models"
)

type ClientUserService interface {
	GetUser() models.GophUser
	SaveUser(user models.GophUser) error
}

// Ensure common implements interface
var _ ClientUserService = (*userService)(nil)

type userService struct {
	user models.GophUser
}

func (s userService) GetUser() models.GophUser {
	return s.user
}

func (s *userService) SaveUser(user models.GophUser) error {
	s.user = user
	return nil
}

func NewUserService() ClientUserService {
	return &userService{}
}
