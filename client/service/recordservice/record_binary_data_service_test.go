package recordservice

import (
	"context"
	"github.com/djokcik/gophkeeper/client/clientmodels"
	"github.com/djokcik/gophkeeper/client/service/mocks"
	mocks2 "github.com/djokcik/gophkeeper/client/service/recordservice/mocks"
	"github.com/djokcik/gophkeeper/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_recordBinaryDataService_RemoveRecordByKey(t *testing.T) {
	t.Run("record binary data should be removed by key", func(t *testing.T) {
		ctx := context.TODO()

		mUser := mocks.ClientUserService{Mock: mock.Mock{}}
		mUser.EXPECT().GetUser().Return(models.ClientUser{Token: "userToken"})

		mAPI := mocks.ClientRPCService{Mock: mock.Mock{}}
		mAPI.EXPECT().RemoveRecordBinaryDataByKey(ctx, "userToken", "testKey").Return(nil)

		svc := recordBinaryDataService{user: &mUser, api: &mAPI}
		err := svc.RemoveRecordByKey(ctx, "testKey")

		require.Equal(t, err, nil)
	})
}

func Test_recordBinaryDataService_LoadRecordByKey(t *testing.T) {
	t.Run("record binary data should be returned by key", func(t *testing.T) {
		ctx := context.TODO()

		user := models.ClientUser{Token: "userToken"}

		mUser := mocks.ClientUserService{Mock: mock.Mock{}}
		mUser.EXPECT().GetUser().Return(user)

		mAPI := mocks.ClientRPCService{Mock: mock.Mock{}}
		mAPI.EXPECT().LoadRecordBinaryDataByKey(ctx, "userToken", "testKey").Return("testData", nil)

		mRecord := mocks2.ClientRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().LoadRecordByKey(ctx, user, mock.Anything, mock.Anything).
			Run(func(ctx context.Context, user models.ClientUser, response interface{}, loadFn func() (string, error)) {
				value := response.(*clientmodels.RecordBinaryData)
				*value = clientmodels.RecordBinaryData{Data: []byte("testText"), Comment: "testComment"}

				data, err := loadFn()
				require.Equal(t, err, nil)
				require.Equal(t, data, "testData")
			}).Return(nil)

		svc := recordBinaryDataService{api: &mAPI, record: &mRecord, user: &mUser}
		binaryData, err := svc.LoadRecordByKey(ctx, "testKey")

		require.Equal(t, err, nil)
		require.Equal(t, binaryData, clientmodels.RecordBinaryData{Data: []byte("testText"), Comment: "testComment"})
	})
}

func Test_recordBinaryDataService_SaveRecord(t *testing.T) {
	t.Run("record binary data should be saved", func(t *testing.T) {
		ctx := context.TODO()

		user := models.ClientUser{Token: "userToken"}
		data := clientmodels.RecordBinaryData{Data: []byte("testText")}

		mUser := mocks.ClientUserService{Mock: mock.Mock{}}
		mUser.EXPECT().GetUser().Return(user)

		mAPI := mocks.ClientRPCService{Mock: mock.Mock{}}
		mAPI.EXPECT().SaveRecordBinaryData(ctx, "userToken", "testKey", "encryptedData").Return(nil)

		mRecord := mocks2.ClientRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().SaveRecord(ctx, user, data, mock.Anything).
			Run(func(ctx context.Context, user models.ClientUser, data interface{}, updateFn func(string) error) {
				updateFn("encryptedData")
			}).Return(nil)

		service := recordBinaryDataService{api: &mAPI, user: &mUser, record: &mRecord}
		err := service.SaveRecord(ctx, "testKey", clientmodels.RecordBinaryData{Data: []byte("testText")})

		require.Equal(t, err, nil)
	})
}
