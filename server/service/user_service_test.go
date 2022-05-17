package service

import (
	"context"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/pkg/common/mocks"
	"github.com/djokcik/gophkeeper/server"
	mocks2 "github.com/djokcik/gophkeeper/server/storage/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_userService_GenerateToken(t *testing.T) {
	t.Run("should generated token", func(t *testing.T) {
		m := mocks.AuthUtilsService{Mock: mock.Mock{}}
		m.EXPECT().CreateToken("key", "username").Return("secretToken", nil)

		service := userService{auth: &m, cfg: server.Config{JWTSecretKey: "key"}}

		token, err := service.GenerateToken(context.Background(), models.User{Username: "username"})

		m.AssertNumberOfCalls(t, "CreateToken", 1)
		require.Equal(t, err, nil)
		require.Equal(t, token, "secretToken")
	})
}

func Test_userService_GetUserByUsername(t *testing.T) {
	t.Run("should return user by username", func(t *testing.T) {
		ctx := context.TODO()

		m := mocks2.Storage{Mock: mock.Mock{}}
		m.On("UserByUsername", ctx, "my_user").Return(models.User{Username: "user"}, nil)
		m.EXPECT().UserByUsername(ctx, "my_user").Return(models.User{Username: "user"}, nil)

		service := userService{repo: &m}
		user, err := service.GetUserByUsername(ctx, "my_user")

		m.AssertNumberOfCalls(t, "UserByUsername", 1)
		require.Equal(t, err, nil)
		require.Equal(t, user, models.User{Username: "user"})
	})
}

func Test_userService_CreateUser(t *testing.T) {
	t.Run("should created user", func(t *testing.T) {
		ctx := context.TODO()

		mAuth := mocks.AuthUtilsService{Mock: mock.Mock{}}
		mAuth.EXPECT().HashAndSalt("my_password", "pepper").Return("testPassword", nil)

		mRepo := mocks2.Storage{Mock: mock.Mock{}}
		mRepo.EXPECT().CreateUser(ctx, models.User{Username: "user", Password: "testPassword"}).Return(nil)

		service := userService{auth: &mAuth, repo: &mRepo, cfg: server.Config{PasswordPepper: "pepper"}}
		err := service.CreateUser(ctx, "user", "my_password")

		mAuth.AssertNumberOfCalls(t, "HashAndSalt", 1)
		mRepo.AssertNumberOfCalls(t, "CreateUser", 1)
		require.Equal(t, err, nil)
	})
}
