package controller

import (
	"gophkeeper/client/registry"
	"gophkeeper/client/view"
	"gophkeeper/client/view/loadpage"
	"gophkeeper/client/view/removepage"
	"gophkeeper/client/view/savepage"
)

type UIController struct {
	RegisterLoginPage *view.AuthPage
	MainPage          *view.MainPage
	RemovePage        *removepage.RemovePage

	SavePage              *savepage.SavePage
	LoginPasswordSavePage *savepage.RecordPersonalDataSavePage
	BinaryDataSavePage    *savepage.BinaryDataSavePage
	TextDataSavePage      *savepage.TextDataSavePage
	CardSavePage          *savepage.BankCardSavePage

	LoadPage              *loadpage.LoadPage
	LoginPasswordLoadPage *loadpage.RecordPersonalDataLoadPage
	TextDataLoadPage      *loadpage.TextDataLoadPage
	BinaryDataLoadPage    *loadpage.BinaryDataLoadPage
	BankCardLoadPage      *loadpage.BankCardLoadPage
}

func NewUIController(serviceRegistry registry.ClientServiceRegistry) *UIController {
	return &UIController{
		RegisterLoginPage: view.NewLoginRegisterPage(serviceRegistry),
		MainPage:          view.NewMainPage(serviceRegistry),
		RemovePage:        removepage.NewRemovePage(serviceRegistry),

		SavePage:              savepage.NewSavePage(),
		LoginPasswordSavePage: savepage.NewRecordPersonalDataSavePage(serviceRegistry),
		TextDataSavePage:      savepage.NewTextDataSavePage(serviceRegistry),
		BinaryDataSavePage:    savepage.NewBinaryDataSavePage(),
		CardSavePage:          savepage.NewBankCardSavePage(serviceRegistry),

		LoadPage:              loadpage.NewLoadPage(),
		LoginPasswordLoadPage: loadpage.NewRecordPersonalDataLoadPage(serviceRegistry),
		TextDataLoadPage:      loadpage.NewTextDataLoadPage(serviceRegistry),
		BinaryDataLoadPage:    loadpage.NewBinaryDataLoadPage(),
		BankCardLoadPage:      loadpage.NewBankCardLoadPage(serviceRegistry),
	}
}
