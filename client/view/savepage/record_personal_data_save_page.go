package savepage

import (
	"context"
	"fmt"
	"github.com/djokcik/gophkeeper/client/clientmodels"
	"github.com/djokcik/gophkeeper/client/registry"
	"github.com/djokcik/gophkeeper/client/view"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/marcusolsson/tui-go"
)

// RecordPersonalDataSavePage is widget for Save PersonalData
type RecordPersonalDataSavePage struct {
	view.PageHooks
	serviceRegistry registry.ClientServiceRegistry

	Root *tui.Box

	keyField      *tui.Entry
	loginField    *tui.Entry
	passwordField *tui.Entry
	urlField      *tui.Entry
	commentField  *tui.Entry
	status        *tui.StatusBar

	Submit *tui.Button
	Back   *tui.Button
}

// GetFocusChain returns list of focused widgets
func (p RecordPersonalDataSavePage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.loginField, p.passwordField, p.urlField, p.commentField, p.Submit, p.Back}
}

// GetRoot return Root winget element
func (p RecordPersonalDataSavePage) GetRoot() tui.Widget {
	return p.Root
}

// OnActivated call one time. Needed for navigate between pages
func (p RecordPersonalDataSavePage) OnActivated(fn func(b *tui.Button)) {
	recordPersonalDataService := p.serviceRegistry.GetRecordPersonalDataService()

	p.Back.OnActivated(fn)
	p.Submit.OnActivated(func(b *tui.Button) {
		ctx, log := logging.GetCtxFileLogger(context.Background())

		if p.keyField.Text() == "" || p.loginField.Text() == "" || p.passwordField.Text() == "" {
			p.status.SetText("Не все поля заполнены")
			return
		}

		data := clientmodels.RecordPersonalData{
			Username: p.loginField.Text(),
			Password: p.passwordField.Text(),
			URL:      p.urlField.Text(),
			Comment:  p.commentField.Text(),
		}

		err := recordPersonalDataService.SaveRecord(ctx, p.keyField.Text(), data)
		if err != nil {
			p.status.SetText(fmt.Sprintf("Error: %s", err.Error()))
			log.Error().Err(err).Msg("Submit: invalid SaveRecord")

			return
		}

		p.status.SetText("Данные успешно сохранены")
	})
}

func NewRecordPersonalDataSavePage(serviceRegistry registry.ClientServiceRegistry) *RecordPersonalDataSavePage {
	p := &RecordPersonalDataSavePage{
		Back:            view.NewBackButton(),
		Submit:          view.NewSaveButton(),
		status:          view.NewStatusLabel(),
		serviceRegistry: serviceRegistry,
	}

	window := view.NewWindowBlock("Сохранить запись по ключу")

	p.keyField = view.NewEditBlockWithWindow("Ключ", window)
	p.keyField.SetFocused(true)

	p.loginField = view.NewEditBlockWithWindow("Логин", window)

	p.passwordField = view.NewEditBlockWithWindow("Пароль", window)
	p.passwordField.SetEchoMode(tui.EchoModePassword)

	p.urlField = view.NewEditBlockWithWindow("URL-адрес", window)
	p.commentField = view.NewEditBlockWithWindow("Примечания", window)

	window.Append(view.NewButtons(p.Back, p.Submit))
	p.Root = tui.NewVBox(view.NewContent(window), p.status)

	return p
}
