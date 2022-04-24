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

	m[ctr.SavePage.Back] = ctr.MainPage

	m[ctr.LoadPage.Back] = ctr.MainPage

	return m
}
