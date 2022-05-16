package storage

import (
	"context"
	"errors"
	"gophkeeper/models"
)

//go:generate mockery --name=Storage --with-expecter
type Storage interface {
	CreateUser(ctx context.Context, user models.User) error
	UserByUsername(ctx context.Context, username string) (models.User, error)

	Read(ctx context.Context, username string) (models.StorageData, error)
	Save(ctx context.Context, data models.StorageData) error
}

var (
	ErrNotFound           = errors.New("storage: not found")
	ErrLoginAlreadyExists = errors.New("storage: login already exists")

	ErrInvalidFileVersion  = errors.New("storage: invalid file version")
	ErrNotFoundFileDecrypt = errors.New("storage: not found file decrypted")
)
