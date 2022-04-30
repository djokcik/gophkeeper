package controller

import (
	"gophkeeper/client/registry"
	"gophkeeper/client/view"
	"gophkeeper/client/view/loadpage"
	"gophkeeper/client/view/savepage"
)

type UIController struct {
	RegisterLoginPage *view.LoginRegisterPage
	MainPage          *view.MainPage

	SavePage              *savepage.SavePage
	LoginPasswordSavePage *savepage.RecordPersonalDataSavePage
	BinaryDataSavePage    *savepage.BinaryDataSavePage
	TextDataSavePage      *savepage.TextDataSavePage
	CardSavePage          *savepage.CardSavePage

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

		SavePage:              savepage.NewSavePage(),
		LoginPasswordSavePage: savepage.NewRecordPersonalDataSavePage(serviceRegistry),
		TextDataSavePage:      savepage.NewTextDataSavePage(),
		BinaryDataSavePage:    savepage.NewBinaryDataSavePage(),
		CardSavePage:          savepage.NewCardSavePage(),

		LoadPage:              loadpage.NewLoadPage(),
		LoginPasswordLoadPage: loadpage.NewRecordPersonalDataLoadPage(serviceRegistry),
		TextDataLoadPage:      loadpage.NewTextDataLoadPage(),
		BinaryDataLoadPage:    loadpage.NewBinaryDataLoadPage(),
		BankCardLoadPage:      loadpage.NewBankCardLoadPage(),
	}
}
