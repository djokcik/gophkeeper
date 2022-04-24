package repo

import "gophkeeper/client/repo/models"

type Storage interface {
	GetUser() models.GophUser
	SaveUser(user models.GophUser) error
}

// Ensure service implements interface
var _ Storage = (*InMemoryStorage)(nil)

type InMemoryStorage struct {
	user models.GophUser
}

func (s InMemoryStorage) GetUser() models.GophUser {
	return s.user
}

func (s *InMemoryStorage) SaveUser(user models.GophUser) error {
	s.user = user
	return nil
}

func NewInMemoryStorage() Storage {
	return &InMemoryStorage{}
}
