package service

import (
	"gophkeeper/models"
)

type ClientUserService interface {
	GetUser() models.ClientUser
	SaveUser(user models.ClientUser) error
}

// Ensure common implements interface
var _ ClientUserService = (*userService)(nil)

type userService struct {
	user models.ClientUser
}

func (s userService) GetUser() models.ClientUser {
	return s.user
}

func (s *userService) SaveUser(user models.ClientUser) error {
	s.user = user
	return nil
}

func NewUserService() ClientUserService {
	return &userService{}
}
