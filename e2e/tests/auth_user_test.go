package tests

import (
	"context"
	"github.com/djokcik/gophkeeper/e2e/testhelpers"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"time"
)

type TestGophkeeperAuth struct {
	suite.Suite

	ctx    context.Context
	cancel context.CancelFunc

	app *testhelpers.Application
}

func (suite *TestGophkeeperAuth) SetupSuite() {
	suite.ctx, suite.cancel = testhelpers.CreateContext()
	suite.app = testhelpers.CreateApplication(suite.ctx, suite.T())

	go suite.app.Run()
	time.Sleep(time.Second)
	suite.app.InitClient()
}

func (suite *TestGophkeeperAuth) TearDownSuite() {
	suite.app.Close()
	suite.cancel()
}

func (suite *TestGophkeeperAuth) TestUserAuth() {
	suite.Run("register user", func() {
		token := suite.app.AppRequest.Register()
		require.NotEqual(suite.T(), token, "")
	})

	suite.Run("login user", func() {
		token := suite.app.AppRequest.Login()
		require.NotEqual(suite.T(), token, "")
	})
}
