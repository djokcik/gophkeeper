package common

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gophkeeper/models"
	"gophkeeper/pkg/common/mocks"
	mocks2 "gophkeeper/server/storage/mocks"
	"testing"
)

func Test_authService_GetUserByToken(t *testing.T) {
	t.Run("should return error when user is empty", func(t *testing.T) {
		ctx := context.TODO()
		service := authService{}

		_, err := service.GetUserByToken(ctx, "")

		require.Equal(t, err, ErrUnauthorized)
	})

	t.Run("should be created user", func(t *testing.T) {
		ctx := context.TODO()

		mAuth := mocks.AuthUtilsService{Mock: mock.Mock{}}
		mAuth.EXPECT().ParseToken("testToken", "key").Return("user", nil).Times(1)

		mRepo := mocks2.Storage{Mock: mock.Mock{}}
		mRepo.EXPECT().UserByUsername(ctx, "user").Return(models.User{Username: "test"}, nil).Times(1)

		service := authService{secretKey: "key", authUtils: &mAuth, storage: &mRepo}
		user, err := service.GetUserByToken(ctx, "testToken")

		require.Equal(t, err, nil)
		require.Equal(t, user, models.User{Username: "test"})
	})
}
