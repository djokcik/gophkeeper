package service

import "errors"

var (
	ErrDuplicateName   = errors.New("LoginService: duplicate name")
	ErrInvalidPassword = errors.New("LoginService: invalid username or password")
)
