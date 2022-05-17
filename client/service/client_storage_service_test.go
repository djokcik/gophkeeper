package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gophkeeper/client/clientmodels"
	mocks2 "gophkeeper/client/service/mocks"
	"gophkeeper/client/storage/mocks"
	"gophkeeper/models"
	"gophkeeper/models/rpcdto"
	"testing"
)

func Test_storageService_LoadRecords(t *testing.T) {
	t.Run("records should be returned", func(t *testing.T) {
		ctx := context.TODO()

		mStore := mocks.ClientLocalStorage{}
		mStore.EXPECT().LoadRecords(ctx).Return([]clientmodels.RecordFileLine{{Key: "testKey"}}, nil)

		svc := storageService{store: &mStore}
		records, err := svc.LoadRecords(ctx)

		require.Equal(t, err, nil)
		require.Equal(t, records, []clientmodels.RecordFileLine{{Key: "testKey"}})
	})
}

func Test_storageService_SyncServer(t *testing.T) {
	t.Run("should return error when token isn`t set", func(t *testing.T) {
		ctx := context.TODO()

		mStore := mocks.ClientLocalStorage{}
		mStore.EXPECT().LoadRecords(ctx).Return([]clientmodels.RecordFileLine{
			{Key: "testKeyAction", ActionType: clientmodels.SaveMethod, Method: "actionMethod", Data: "testActionData"},
		}, nil)

		mUser := mocks2.ClientUserService{Mock: mock.Mock{}}
		mUser.EXPECT().GetUser().Return(models.ClientUser{Token: ""})

		svc := storageService{store: &mStore, user: &mUser}
		err := svc.SyncServer(ctx)

		require.Equal(t, err, ErrUnableConnectServer)
	})

	t.Run("client saved actions should be synced to server", func(t *testing.T) {
		ctx := context.TODO()

		mStore := mocks.ClientLocalStorage{}
		mStore.EXPECT().LoadRecords(ctx).Return([]clientmodels.RecordFileLine{
			{Key: "testKeyAction", ActionType: clientmodels.SaveMethod, Method: "actionMethod", Data: "testActionData"},
		}, nil)
		mStore.EXPECT().ClearActions(ctx).Return(nil)

		mUser := mocks2.ClientUserService{Mock: mock.Mock{}}
		mUser.EXPECT().GetUser().Return(models.ClientUser{Token: "userToken"})

		saveRecordDto := rpcdto.SaveRecordRequestDto{Key: "testKeyAction", Data: "testActionData", Token: "userToken"}

		mApi := mocks2.ClientRpcService{Mock: mock.Mock{}}
		mApi.EXPECT().Call(ctx, "actionMethod", saveRecordDto, mock.Anything).Return(nil)

		svc := storageService{api: &mApi, user: &mUser, store: &mStore}
		err := svc.SyncServer(ctx)

		require.Equal(t, err, nil)
		mStore.AssertNumberOfCalls(t, "LoadRecords", 1)
		mStore.AssertNumberOfCalls(t, "ClearActions", 1)
		mUser.AssertNumberOfCalls(t, "GetUser", 1)
		mApi.AssertNumberOfCalls(t, "Call", 1)
	})

	t.Run("client removed actions should be synced to server", func(t *testing.T) {
		ctx := context.TODO()

		mStore := mocks.ClientLocalStorage{}
		mStore.EXPECT().LoadRecords(ctx).Return([]clientmodels.RecordFileLine{
			{Key: "testKeyAction", ActionType: clientmodels.RemoveMethod, Method: "actionMethod"},
		}, nil)
		mStore.EXPECT().ClearActions(ctx).Return(nil)

		mUser := mocks2.ClientUserService{Mock: mock.Mock{}}
		mUser.EXPECT().GetUser().Return(models.ClientUser{Token: "userToken"})

		removeRecordDto := rpcdto.RemoveRecordRequestDto{Key: "testKeyAction", Token: "userToken"}

		mApi := mocks2.ClientRpcService{Mock: mock.Mock{}}
		mApi.EXPECT().Call(ctx, "actionMethod", removeRecordDto, mock.Anything).Return(nil)

		svc := storageService{api: &mApi, user: &mUser, store: &mStore}
		err := svc.SyncServer(ctx)

		require.Equal(t, err, nil)
		mStore.AssertNumberOfCalls(t, "LoadRecords", 1)
		mStore.AssertNumberOfCalls(t, "ClearActions", 1)
		mUser.AssertNumberOfCalls(t, "GetUser", 1)
		mApi.AssertNumberOfCalls(t, "Call", 1)
	})
}
