package controller

import (
	"gophkeeper/client/view"
)

type UIController struct {
	RegisterLoginPage *view.LoginRegisterPage
	MainPage          *view.MainPage
	SavePage          *view.SagePage
	LoadPage          *view.LoadPage
}

func NewUIController() *UIController {
	return &UIController{
		RegisterLoginPage: view.NewLoginRegisterPage(),
		MainPage:          view.NewMainPage(),
		SavePage:          view.NewSagePage(),
		LoadPage:          view.NewLoadPage(),
	}
}
