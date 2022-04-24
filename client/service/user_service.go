package service

import (
	"gophkeeper/models"
)

type UserService interface {
	GetUser() models.GophUser
	SaveUser(user models.GophUser) error
}

// Ensure service implements interface
var _ UserService = (*userService)(nil)

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

func NewUserService() UserService {
	return &userService{}
}
