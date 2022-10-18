package service

import (
	"context"
	"github.com/djokcik/gophkeeper/client/service/mocks"
	"github.com/djokcik/gophkeeper/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_authService_SignIn(t *testing.T) {
	t.Run("user should call login", func(t *testing.T) {
		ctx := context.TODO()

		mAPI := mocks.ClientRPCService{Mock: mock.Mock{}}
		mAPI.EXPECT().Login(ctx, "username", "password").Return("userToken", nil)

		svc := authService{api: &mAPI}
		user, err := svc.SignIn(ctx, "username", "password")

		require.Equal(t, err, nil)
		require.Equal(t, user, models.ClientUser{Username: "username", Password: "password", Token: "userToken"})
	})

	t.Run("user should be returned when api return offline mode", func(t *testing.T) {
		ctx := context.TODO()

		mAPI := mocks.ClientRPCService{Mock: mock.Mock{}}
		mAPI.EXPECT().Login(ctx, "username", "password").Return("", ErrAnonymousUser)

		svc := authService{api: &mAPI}
		user, err := svc.SignIn(ctx, "username", "password")

		require.Equal(t, err, nil)
		require.Equal(t, user, models.ClientUser{Username: "username", Password: "password"})
	})
}

func Test_authService_Register(t *testing.T) {
	t.Run("user should call register", func(t *testing.T) {
		ctx := context.TODO()

		mAPI := mocks.ClientRPCService{Mock: mock.Mock{}}
		mAPI.EXPECT().Register(ctx, "username", "password").Return("userToken", nil)

		svc := authService{api: &mAPI}
		user, err := svc.Register(ctx, "username", "password")

		require.Equal(t, err, nil)
		require.Equal(t, user, models.ClientUser{Username: "username", Password: "password", Token: "userToken"})
	})

	t.Run("user should be returned when api return offline mode", func(t *testing.T) {
		ctx := context.TODO()

		mAPI := mocks.ClientRPCService{Mock: mock.Mock{}}
		mAPI.EXPECT().Register(ctx, "username", "password").Return("", ErrAnonymousUser)

		svc := authService{api: &mAPI}
		user, err := svc.Register(ctx, "username", "password")

		require.Equal(t, err, nil)
		require.Equal(t, user, models.ClientUser{Username: "username", Password: "password"})
	})
}
