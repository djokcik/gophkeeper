package rpchandler

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gophkeeper/models"
	"gophkeeper/models/rpcdto"
	"gophkeeper/pkg/common/mocks"
	mocks2 "gophkeeper/server/service/mocks"
	"testing"
)

func TestRpcHandler_LoadRecordTextDataByKeyHandler(t *testing.T) {
	t.Run("text data should be returned", func(t *testing.T) {
		ctx := mock.Anything

		mAuth := mocks.AuthService{Mock: mock.Mock{}}
		mAuth.EXPECT().GetUserByToken(ctx, "token").Return(models.User{Username: "test"}, nil)

		mRecord := mocks2.ServerRecordTextDataService{Mock: mock.Mock{}}
		mRecord.EXPECT().Load(ctx, "testKey", "test").Return("testData", nil)

		var reply string

		handler := &RpcHandler{auth: &mAuth, recordTextData: &mRecord}
		err := handler.LoadRecordTextDataByKeyHandler(rpcdto.LoadRecordRequestDto{Key: "testKey", Token: "token"}, &reply)

		require.Equal(t, err, nil)
		require.Equal(t, reply, "testData")
		mAuth.AssertNumberOfCalls(t, "GetUserByToken", 1)
		mRecord.AssertNumberOfCalls(t, "Load", 1)
	})
}

func TestRpcHandler_SaveRecordTextDataHandler(t *testing.T) {
	t.Run("text data should be saved", func(t *testing.T) {
		ctx := mock.Anything

		mAuth := mocks.AuthService{Mock: mock.Mock{}}
		mAuth.EXPECT().GetUserByToken(ctx, "token").Return(models.User{Username: "test"}, nil)

		mRecord := mocks2.ServerRecordTextDataService{Mock: mock.Mock{}}
		mRecord.EXPECT().Save(ctx, "testKey", "test", "testData").Return(nil)

		var reply struct{}

		handler := &RpcHandler{auth: &mAuth, recordTextData: &mRecord}
		err := handler.SaveRecordTextDataHandler(rpcdto.SaveRecordRequestDto{Key: "testKey", Token: "token", Data: "testData"}, &reply)

		require.Equal(t, err, nil)
		mAuth.AssertNumberOfCalls(t, "GetUserByToken", 1)
		mRecord.AssertNumberOfCalls(t, "Save", 1)
	})
}

func TestRpcHandler_RemoveRecordTextDataByKeyHandler(t *testing.T) {
	t.Run("text data should be removed", func(t *testing.T) {
		ctx := mock.Anything

		mAuth := mocks.AuthService{Mock: mock.Mock{}}
		mAuth.EXPECT().GetUserByToken(ctx, "token").Return(models.User{Username: "test"}, nil)

		mRecord := mocks2.ServerRecordTextDataService{Mock: mock.Mock{}}
		mRecord.EXPECT().Remove(ctx, "testKey", "test").Return(nil)

		var reply struct{}

		handler := &RpcHandler{auth: &mAuth, recordTextData: &mRecord}
		err := handler.RemoveRecordTextDataByKeyHandler(rpcdto.RemoveRecordRequestDto{Key: "testKey", Token: "token"}, &reply)

		require.Equal(t, err, nil)
		mAuth.AssertNumberOfCalls(t, "GetUserByToken", 1)
		mRecord.AssertNumberOfCalls(t, "Remove", 1)
	})
}
