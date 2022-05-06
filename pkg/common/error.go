package common

import "errors"

var (
	ErrUnauthorized       = errors.New("unauthorized")
	ErrWrongPassword      = errors.New("authenticate: invalid password")
	ErrNotAuthenticated   = errors.New("service: no authenticted user found in the context")
	ErrInvalidAccessToken = errors.New("invalid auth token")
)
