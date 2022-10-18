package rpchandler

import (
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/models/rpcdto"
	"github.com/djokcik/gophkeeper/pkg/common/mocks"
	mocks2 "github.com/djokcik/gophkeeper/server/service/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRpcHandler_LoadRecordBankCardDataByKeyHandler(t *testing.T) {
	t.Run("bank card data should be returned", func(t *testing.T) {
		ctx := mock.Anything

		mAuth := mocks.AuthService{Mock: mock.Mock{}}
		mAuth.EXPECT().GetUserByToken(ctx, "token").Return(models.User{Username: "test"}, nil)

		mRecord := mocks2.ServerRecordBankCardDataService{Mock: mock.Mock{}}
		mRecord.EXPECT().Load(ctx, "testKey", "test").Return("testData", nil)

		var reply string

		handler := &RPCHandler{auth: &mAuth, recordBankCard: &mRecord}
		err := handler.LoadRecordBankCardByKeyHandler(rpcdto.LoadRecordRequestDto{Key: "testKey", Token: "token"}, &reply)

		require.Equal(t, err, nil)
		require.Equal(t, reply, "testData")
		mAuth.AssertNumberOfCalls(t, "GetUserByToken", 1)
		mRecord.AssertNumberOfCalls(t, "Load", 1)
	})
}

func TestRpcHandler_SaveRecordBankCardHandler(t *testing.T) {
	t.Run("bank card data should be saved", func(t *testing.T) {
		ctx := mock.Anything

		mAuth := mocks.AuthService{Mock: mock.Mock{}}
		mAuth.EXPECT().GetUserByToken(ctx, "token").Return(models.User{Username: "test"}, nil)

		mRecord := mocks2.ServerRecordBankCardDataService{Mock: mock.Mock{}}
		mRecord.EXPECT().Save(ctx, "testKey", "test", "testData").Return(nil)

		var reply struct{}

		handler := &RPCHandler{auth: &mAuth, recordBankCard: &mRecord}
		err := handler.SaveRecordBankCardHandler(rpcdto.SaveRecordRequestDto{Key: "testKey", Token: "token", Data: "testData"}, &reply)

		require.Equal(t, err, nil)
		mAuth.AssertNumberOfCalls(t, "GetUserByToken", 1)
		mRecord.AssertNumberOfCalls(t, "Save", 1)
	})
}

func TestRpcHandler_RemoveRecordBankCardDataByKeyHandler(t *testing.T) {
	t.Run("bank card data should be removed", func(t *testing.T) {
		ctx := mock.Anything

		mAuth := mocks.AuthService{Mock: mock.Mock{}}
		mAuth.EXPECT().GetUserByToken(ctx, "token").Return(models.User{Username: "test"}, nil)

		mRecord := mocks2.ServerRecordBankCardDataService{Mock: mock.Mock{}}
		mRecord.EXPECT().Remove(ctx, "testKey", "test").Return(nil)

		var reply struct{}

		handler := &RPCHandler{auth: &mAuth, recordBankCard: &mRecord}
		err := handler.RemoveRecordBankCardByKeyHandler(rpcdto.RemoveRecordRequestDto{Key: "testKey", Token: "token"}, &reply)

		require.Equal(t, err, nil)
		mAuth.AssertNumberOfCalls(t, "GetUserByToken", 1)
		mRecord.AssertNumberOfCalls(t, "Remove", 1)
	})
}
