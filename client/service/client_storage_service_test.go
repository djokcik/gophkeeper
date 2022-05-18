package service

import (
	"context"
	"github.com/djokcik/gophkeeper/client/clientmodels"
	mocks2 "github.com/djokcik/gophkeeper/client/service/mocks"
	"github.com/djokcik/gophkeeper/client/storage/mocks"
	"github.com/djokcik/gophkeeper/models"
	"github.com/djokcik/gophkeeper/models/rpcdto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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

		mAPI := mocks2.ClientRPCService{Mock: mock.Mock{}}
		mAPI.EXPECT().Call(ctx, "actionMethod", saveRecordDto, mock.Anything).Return(nil)

		svc := storageService{api: &mAPI, user: &mUser, store: &mStore}
		err := svc.SyncServer(ctx)

		require.Equal(t, err, nil)
		mStore.AssertNumberOfCalls(t, "LoadRecords", 1)
		mStore.AssertNumberOfCalls(t, "ClearActions", 1)
		mUser.AssertNumberOfCalls(t, "GetUser", 1)
		mAPI.AssertNumberOfCalls(t, "Call", 1)
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

		mAPI := mocks2.ClientRPCService{Mock: mock.Mock{}}
		mAPI.EXPECT().Call(ctx, "actionMethod", removeRecordDto, mock.Anything).Return(nil)

		svc := storageService{api: &mAPI, user: &mUser, store: &mStore}
		err := svc.SyncServer(ctx)

		require.Equal(t, err, nil)
		mStore.AssertNumberOfCalls(t, "LoadRecords", 1)
		mStore.AssertNumberOfCalls(t, "ClearActions", 1)
		mUser.AssertNumberOfCalls(t, "GetUser", 1)
		mAPI.AssertNumberOfCalls(t, "Call", 1)
	})
}
