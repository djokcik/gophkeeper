package service

import "errors"

var (
	ErrDuplicateName   = errors.New("AuthService: duplicate name")
	ErrInvalidPassword = errors.New("AuthService: invalid username or password")

	ErrUnableConnectServer = errors.New("RpcService: unable connect to server")
)
