package service

import (
	"github.com/djokcik/gophkeeper/models"
)

//go:generate mockery --name=ClientUserService --with-expecter

// ClientUserService provides methods for save and get users
type ClientUserService interface {
	GetUser() models.ClientUser
	SaveUser(user models.ClientUser) error
}

// Ensure common implements interface
var _ ClientUserService = (*userService)(nil)

type userService struct {
	user models.ClientUser
}

// GetUser returns user
func (s userService) GetUser() models.ClientUser {
	return s.user
}

// SaveUser save user in memory
func (s *userService) SaveUser(user models.ClientUser) error {
	s.user = user
	return nil
}

func NewUserService() ClientUserService {
	return &userService{}
}
