package service

import "errors"

var (
	ErrUnauthorized  = errors.New("unauthorized")
	ErrWrongPassword = errors.New("authenticate: invalid password")

	ErrNotFoundRecord = errors.New("service: not found record")
)
