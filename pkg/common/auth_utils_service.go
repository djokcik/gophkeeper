package common

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gophkeeper/models"
	"time"
)

//go:generate mockery --name=AuthUtilsService --with-expecter
type AuthUtilsService interface {
	CreateToken(secretKey string, username string) (string, error)
	ParseToken(accessToken string, secretKey string) (string, error)
	HashAndSalt(pwd string, pepper string) (string, error)
	CompareHashAndPassword(password string, hash string) error
}

func NewAuthUtilsService() AuthUtilsService {
	return &userUtilsService{}
}

type userUtilsService struct {
}

func (a userUtilsService) CompareHashAndPassword(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (a userUtilsService) HashAndSalt(pwd string, pepper string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd+pepper), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcryptPassword: %w", err)
	}

	return string(hash), nil
}

func (a userUtilsService) CreateToken(secretKey string, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    "gophkeeper",
		},
	})

	return token.SignedString([]byte(secretKey))
}

func (a userUtilsService) ParseToken(accessToken string, secretKey string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", ErrInvalidAccessToken
}
