package controller

import (
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/view"
)

type Route map[*tui.Button]view.Page

func GenerateRouter(ctr *UIController) Route {
	m := make(Route)
	m[ctr.RegisterLoginPage.Login] = ctr.MainPage
	m[ctr.RegisterLoginPage.Register] = ctr.MainPage

	m[ctr.MainPage.SaveData] = ctr.SavePage
	m[ctr.MainPage.LoadData] = ctr.LoadPage

	m[ctr.SavePage.LoginPassword] = ctr.LoginPasswordSavePage
	m[ctr.SavePage.BinButton] = ctr.BinaryDataSavePage
	m[ctr.SavePage.TextButton] = ctr.TextDataSavePage
	m[ctr.SavePage.CardButton] = ctr.CardSavePage
	m[ctr.SavePage.Back] = ctr.MainPage

	m[ctr.LoadPage.Back] = ctr.MainPage
	m[ctr.LoadPage.LoginPassword] = ctr.LoginPasswordLoadPage
	m[ctr.LoadPage.TextButton] = ctr.TextDataLoadPage
	m[ctr.LoadPage.BinButton] = ctr.BinaryDataLoadPage
	m[ctr.LoadPage.CardButton] = ctr.BankCardLoadPage

	m[ctr.LoginPasswordLoadPage.Back] = ctr.LoadPage
	m[ctr.TextDataLoadPage.Back] = ctr.LoadPage
	m[ctr.BinaryDataLoadPage.Back] = ctr.LoadPage
	m[ctr.BankCardLoadPage.Back] = ctr.LoadPage

	m[ctr.LoginPasswordSavePage.Back] = ctr.SavePage
	m[ctr.BinaryDataSavePage.Back] = ctr.SavePage
	m[ctr.TextDataSavePage.Back] = ctr.SavePage
	m[ctr.CardSavePage.Back] = ctr.SavePage

	return m
}
