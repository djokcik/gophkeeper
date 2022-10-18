package models

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

type (
	Claims struct {
		jwt.StandardClaims
		Username string
	}

	ClientUser struct {
		Username string
		Password string
		Token    string
	}

	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func (u User) Validate() error {
	if len(u.Username) < 3 || len(u.Username) > 20 {
		return ErrUsernameLength
	}

	if u.Password == "" {
		return ErrPasswordEmpty
	}

	if len(u.Password) < 3 || len(u.Password) > 256 {
		return ErrPasswordLength
	}

	return nil
}

var (
	ErrUsernameLength = errors.New("validate username: invalid length")
	ErrPasswordEmpty  = errors.New("validate password: empty")
	ErrPasswordLength = errors.New("validate password: invalid length")
)
