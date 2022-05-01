package service

import "errors"

var (
	ErrDuplicateName   = errors.New("AuthService: duplicate name")
	ErrInvalidPassword = errors.New("AuthService: invalid username or password")

	ErrUnableConnectServer = errors.New("ClientRpcService: unable connect to server")
)
