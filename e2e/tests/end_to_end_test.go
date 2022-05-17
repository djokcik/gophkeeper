package tests

import (
	"context"
	"github.com/djokcik/gophkeeper/e2e/testhelpers"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"time"
)

type TestEndToEnd struct {
	suite.Suite

	ctx    context.Context
	cancel context.CancelFunc

	app   *testhelpers.Application
	token string
}

func (suite *TestEndToEnd) SetupSuite() {
	suite.ctx, suite.cancel = testhelpers.CreateContext()
	suite.app = testhelpers.CreateApplication(suite.ctx, suite.T())
	suite.app.ClearUsers()

	go suite.app.Run()
	time.Sleep(time.Second)
	suite.app.InitClient()
	suite.token = suite.app.AppRequest.Register()
}

func (suite *TestEndToEnd) TearDownSuite() {
	suite.app.Close()
	suite.cancel()
}

func (suite *TestEndToEnd) TestRecordsOperations() {
	suite.Run("save personal data", func() {
		suite.app.AppRequest.SaveData(suite.token, "testKey", "testData", testhelpers.PersonalDataRecord)
	})

	suite.Run("load personal data by key", func() {
		data := suite.app.AppRequest.LoadData(suite.token, "testKey", testhelpers.PersonalDataRecord)
		require.Equal(suite.T(), data, "testData")
	})

	suite.Run("remove personal data", func() {
		suite.app.AppRequest.RemoveData(suite.token, "testKey", testhelpers.PersonalDataRecord)
		data := suite.app.AppRequest.LoadData(suite.token, "testKey", testhelpers.PersonalDataRecord)
		require.Equal(suite.T(), data, "")
	})

	suite.Run("save text data", func() {
		suite.app.AppRequest.SaveData(suite.token, "testKey", "testData", testhelpers.TextDataRecord)
	})

	suite.Run("load text data by key", func() {
		data := suite.app.AppRequest.LoadData(suite.token, "testKey", testhelpers.TextDataRecord)
		require.Equal(suite.T(), data, "testData")
	})

	suite.Run("remove text data", func() {
		suite.app.AppRequest.RemoveData(suite.token, "testKey", testhelpers.TextDataRecord)
		data := suite.app.AppRequest.LoadData(suite.token, "testKey", testhelpers.TextDataRecord)
		require.Equal(suite.T(), data, "")
	})

	suite.Run("save binary data", func() {
		suite.app.AppRequest.SaveData(suite.token, "testKey", "testData", testhelpers.BinaryDataRecord)
	})

	suite.Run("load binary data by key", func() {
		data := suite.app.AppRequest.LoadData(suite.token, "testKey", testhelpers.BinaryDataRecord)
		require.Equal(suite.T(), data, "testData")
	})

	suite.Run("remove binary data", func() {
		suite.app.AppRequest.RemoveData(suite.token, "testKey", testhelpers.BinaryDataRecord)
		data := suite.app.AppRequest.LoadData(suite.token, "testKey", testhelpers.BinaryDataRecord)
		require.Equal(suite.T(), data, "")
	})

	suite.Run("save bank card data", func() {
		suite.app.AppRequest.SaveData(suite.token, "testKey", "testData", testhelpers.BankCardRecord)
	})

	suite.Run("load bank card  data by key", func() {
		data := suite.app.AppRequest.LoadData(suite.token, "testKey", testhelpers.BankCardRecord)
		require.Equal(suite.T(), data, "testData")
	})

	suite.Run("remove bank card data", func() {
		suite.app.AppRequest.RemoveData(suite.token, "testKey", testhelpers.BankCardRecord)
		data := suite.app.AppRequest.LoadData(suite.token, "testKey", testhelpers.BankCardRecord)
		require.Equal(suite.T(), data, "")
	})
}

func (suite *TestEndToEnd) TestEndToEnd() {
	suite.Run("update existing record", func() {
		suite.app.AppRequest.SaveData(suite.token, "customKey", "data", testhelpers.TextDataRecord)
		suite.app.AppRequest.SaveData(suite.token, "customKey", "textData", testhelpers.TextDataRecord)
		data := suite.app.AppRequest.LoadData(suite.token, "customKey", testhelpers.TextDataRecord)
		require.Equal(suite.T(), data, "textData")
	})

	suite.Run("save same key in different record", func() {
		suite.app.AppRequest.SaveData(suite.token, "key", "textData", testhelpers.TextDataRecord)
		suite.app.AppRequest.SaveData(suite.token, "key", "personalData", testhelpers.PersonalDataRecord)
		suite.app.AppRequest.SaveData(suite.token, "key", "binaryData", testhelpers.BinaryDataRecord)
		suite.app.AppRequest.SaveData(suite.token, "key", "bankCardData", testhelpers.BankCardRecord)

		var data string

		data = suite.app.AppRequest.LoadData(suite.token, "key", testhelpers.TextDataRecord)
		require.Equal(suite.T(), data, "textData")

		data = suite.app.AppRequest.LoadData(suite.token, "key", testhelpers.PersonalDataRecord)
		require.Equal(suite.T(), data, "personalData")

		data = suite.app.AppRequest.LoadData(suite.token, "key", testhelpers.BinaryDataRecord)
		require.Equal(suite.T(), data, "binaryData")

		data = suite.app.AppRequest.LoadData(suite.token, "key", testhelpers.BankCardRecord)
		require.Equal(suite.T(), data, "bankCardData")
	})
}
