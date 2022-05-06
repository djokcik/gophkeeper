package service

import "errors"

var (
	ErrUnauthorized  = errors.New("unauthorized")
	ErrWrongPassword = errors.New("authenticate: invalid password")
)
