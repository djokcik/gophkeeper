package rpchandler

import (
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/models/rpcdto"
	"github.com/djokcik/gophkeeper/server/service/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRpcHandler_SignInHandler(t *testing.T) {
	t.Run("token should be generated", func(t *testing.T) {
		mUser := mocks.ServerUserService{Mock: mock.Mock{}}
		mUser.EXPECT().Authenticate(mock.Anything, "test_login", "test_password").Return("test_token", nil)

		var reply string

		handler := &RPCHandler{user: &mUser}
		err := handler.SignInHandler(rpcdto.LoginDto{Login: "test_login", Password: "test_password"}, &reply)

		require.Equal(t, err, nil)
		require.Equal(t, reply, "test_token")
	})
}

func TestRpcHandler_RegisterHandler(t *testing.T) {
	t.Run("user should be registered", func(t *testing.T) {
		ctx := mock.Anything

		mUser := mocks.ServerUserService{Mock: mock.Mock{}}
		mUser.EXPECT().CreateUser(ctx, "username", "pass").Return(nil)
		mUser.EXPECT().GetUserByUsername(ctx, "username").Return(models.User{Username: "testUser"}, nil)
		mUser.EXPECT().GenerateToken(ctx, models.User{Username: "testUser"}).Return("token", nil)

		var reply string

		handler := &RPCHandler{user: &mUser}
		err := handler.RegisterHandler(rpcdto.RegisterDto{Login: "username", Password: "pass"}, &reply)

		require.Equal(t, err, nil)
		require.Equal(t, reply, "token")
	})
}
