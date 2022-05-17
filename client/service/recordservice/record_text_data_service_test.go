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

func Test_recordTextDataService_RemoveRecordByKey(t *testing.T) {
	t.Run("record text data should be removed by key", func(t *testing.T) {
		ctx := context.TODO()

		mUser := mocks.ClientUserService{Mock: mock.Mock{}}
		mUser.EXPECT().GetUser().Return(models.ClientUser{Token: "userToken"})

		mApi := mocks.ClientRpcService{Mock: mock.Mock{}}
		mApi.EXPECT().RemoveRecordTextDataByKey(ctx, "userToken", "testKey").Return(nil)

		svc := recordTextDataService{user: &mUser, api: &mApi}
		err := svc.RemoveRecordByKey(ctx, "testKey")

		require.Equal(t, err, nil)
	})
}

func Test_recordTextDataService_LoadRecordByKey(t *testing.T) {
	t.Run("record text data should be returned by key", func(t *testing.T) {
		ctx := context.TODO()

		user := models.ClientUser{Token: "userToken"}

		mUser := mocks.ClientUserService{Mock: mock.Mock{}}
		mUser.EXPECT().GetUser().Return(user)

		mApi := mocks.ClientRpcService{Mock: mock.Mock{}}
		mApi.EXPECT().LoadRecordTextDataByKey(ctx, "userToken", "testKey").Return("testData", nil)

		mRecord := mocks2.ClientRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().LoadRecordByKey(ctx, user, mock.Anything, mock.Anything).
			Run(func(ctx context.Context, user models.ClientUser, response interface{}, loadFn func() (string, error)) {
				value := response.(*clientmodels.RecordTextData)
				*value = clientmodels.RecordTextData{Text: "testText", Comment: "testComment"}

				data, err := loadFn()
				require.Equal(t, err, nil)
				require.Equal(t, data, "testData")
			}).Return(nil)

		svc := recordTextDataService{api: &mApi, record: &mRecord, user: &mUser}
		textData, err := svc.LoadRecordByKey(ctx, "testKey")

		require.Equal(t, err, nil)
		require.Equal(t, textData, clientmodels.RecordTextData{Text: "testText", Comment: "testComment"})
	})
}

func Test_recordTextDataService_SaveRecord(t *testing.T) {
	t.Run("record text data should be saved", func(t *testing.T) {
		ctx := context.TODO()

		user := models.ClientUser{Token: "userToken"}
		data := clientmodels.RecordTextData{Text: "test"}

		mUser := mocks.ClientUserService{Mock: mock.Mock{}}
		mUser.EXPECT().GetUser().Return(user)

		mApi := mocks.ClientRpcService{Mock: mock.Mock{}}
		mApi.EXPECT().SaveRecordTextData(ctx, "userToken", "testKey", "encryptedData").Return(nil)

		mRecord := mocks2.ClientRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().SaveRecord(ctx, user, data, mock.Anything).
			Run(func(ctx context.Context, user models.ClientUser, data interface{}, updateFn func(string) error) {
				updateFn("encryptedData")
			}).Return(nil)

		service := recordTextDataService{api: &mApi, user: &mUser, record: &mRecord}
		err := service.SaveRecord(ctx, "testKey", clientmodels.RecordTextData{Text: "test"})

		require.Equal(t, err, nil)
	})
}
