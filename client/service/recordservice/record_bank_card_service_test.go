package recordservice

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gophkeeper/client/clientmodels"
	"gophkeeper/client/service/mocks"
	mocks2 "gophkeeper/client/service/recordservice/mocks"
	"gophkeeper/models"
	"testing"
)

func Test_recordBankCardDataService_RemoveRecordByKey(t *testing.T) {
	t.Run("record text data should be removed by key", func(t *testing.T) {
		ctx := context.TODO()

		mUser := mocks.ClientUserService{Mock: mock.Mock{}}
		mUser.EXPECT().GetUser().Return(models.ClientUser{Token: "userToken"})

		mApi := mocks.ClientRpcService{Mock: mock.Mock{}}
		mApi.EXPECT().RemoveRecordBankCardByKey(ctx, "userToken", "testKey").Return(nil)

		svc := recordBankCardService{user: &mUser, api: &mApi}
		err := svc.RemoveRecordByKey(ctx, "testKey")

		require.Equal(t, err, nil)
	})
}

func Test_recordBankCardDataService_LoadRecordByKey(t *testing.T) {
	t.Run("record text data should be returned by key", func(t *testing.T) {
		ctx := context.TODO()

		user := models.ClientUser{Token: "userToken"}

		mUser := mocks.ClientUserService{Mock: mock.Mock{}}
		mUser.EXPECT().GetUser().Return(user)

		mApi := mocks.ClientRpcService{Mock: mock.Mock{}}
		mApi.EXPECT().LoadRecordBankCardByKey(ctx, "userToken", "testKey").Return("testData", nil)

		mRecord := mocks2.ClientRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().LoadRecordByKey(ctx, user, mock.Anything, mock.Anything).
			Run(func(ctx context.Context, user models.ClientUser, response interface{}, loadFn func() (string, error)) {
				value := response.(*clientmodels.RecordBankCardData)
				*value = clientmodels.RecordBankCardData{CardNumber: "test", CVV: "cvv", Comment: "testComment"}

				data, err := loadFn()
				require.Equal(t, err, nil)
				require.Equal(t, data, "testData")
			}).Return(nil)

		svc := recordBankCardService{api: &mApi, record: &mRecord, user: &mUser}
		textData, err := svc.LoadRecordByKey(ctx, "testKey")

		require.Equal(t, err, nil)
		require.Equal(t, textData, clientmodels.RecordBankCardData{CardNumber: "test", CVV: "cvv", Comment: "testComment"})
	})
}

func Test_recordBankCardDataService_SaveRecord(t *testing.T) {
	t.Run("record text data should be saved", func(t *testing.T) {
		ctx := context.TODO()

		user := models.ClientUser{Token: "userToken"}
		data := clientmodels.RecordBankCardData{CardNumber: "test", CVV: "cvv"}

		mUser := mocks.ClientUserService{Mock: mock.Mock{}}
		mUser.EXPECT().GetUser().Return(user)

		mApi := mocks.ClientRpcService{Mock: mock.Mock{}}
		mApi.EXPECT().SaveRecordBankCard(ctx, "userToken", "testKey", "encryptedData").Return(nil)

		mRecord := mocks2.ClientRecordService{Mock: mock.Mock{}}
		mRecord.EXPECT().SaveRecord(ctx, user, data, mock.Anything).
			Run(func(ctx context.Context, user models.ClientUser, data interface{}, updateFn func(string) error) {
				updateFn("encryptedData")
			}).Return(nil)

		service := recordBankCardService{api: &mApi, user: &mUser, record: &mRecord}
		err := service.SaveRecord(ctx, "testKey", clientmodels.RecordBankCardData{CardNumber: "test", CVV: "cvv"})

		require.Equal(t, err, nil)
	})
}