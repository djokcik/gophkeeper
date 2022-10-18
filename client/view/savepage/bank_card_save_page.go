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

// BankCardSavePage is widget for Save BankCard
type BankCardSavePage struct {
	view.PageHooks

	serviceRegistry registry.ClientServiceRegistry
	Root            *tui.Box

	keyField        *tui.Entry
	cardNumberField *tui.Entry
	yearField       *tui.Entry
	cvvField        *tui.Entry
	comment         *tui.Entry
	status          *tui.StatusBar

	Submit *tui.Button
	Back   *tui.Button
}

// GetFocusChain returns list of focused widgets
func (p BankCardSavePage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.cardNumberField, p.yearField, p.cvvField, p.comment, p.Submit, p.Back}
}

// GetRoot return Root winget element
func (p BankCardSavePage) GetRoot() tui.Widget {
	return p.Root
}

// OnActivated call one time. Needed for navigate between pages
func (p BankCardSavePage) OnActivated(fn func(b *tui.Button)) {
	recordBankCardService := p.serviceRegistry.GetRecordBankCardService()

	p.Back.OnActivated(fn)
	p.Submit.OnActivated(func(b *tui.Button) {
		ctx := context.Background()
		log := logging.NewFileLogger()

		if p.keyField.Text() == "" || p.cardNumberField.Text() == "" || p.yearField.Text() == "" || p.cvvField.Text() == "" {
			p.status.SetText("Не все поля заполнены")
			return
		}

		data := clientmodels.RecordBankCardData{
			CardNumber: p.cardNumberField.Text(),
			Year:       p.yearField.Text(),
			CVV:        p.cvvField.Text(),
			Comment:    p.comment.Text(),
		}

		err := recordBankCardService.SaveRecord(ctx, p.keyField.Text(), data)
		if err != nil {
			p.status.SetText(fmt.Sprintf("Error: %s", err.Error()))
			log.Error().Err(err).Msg("Submit: invalid SaveRecord")

			return
		}

		p.status.SetText("Банковская карта успешно сохранена")
	})
}

func NewBankCardSavePage(serviceRegistry registry.ClientServiceRegistry) *BankCardSavePage {
	p := &BankCardSavePage{
		Back:            view.NewBackButton(),
		Submit:          view.NewSaveButton(),
		status:          view.NewStatusLabel(),
		serviceRegistry: serviceRegistry,
	}

	window := view.NewWindowBlock("Сохранить запись по ключу")

	p.keyField = view.NewEditBlockWithWindow("Ключ", window)
	p.keyField.SetFocused(true)

	p.cardNumberField = view.NewEditBlockWithWindow("Номер карты", window)
	p.yearField = view.NewEditBlockWithWindow("Год выпуска", window)

	p.cvvField = view.NewEditBlockWithWindow("cvv", window)
	p.cvvField.SetEchoMode(tui.EchoModePassword)

	p.comment = view.NewEditBlockWithWindow("Примечание", window)

	window.Append(view.NewButtons(p.Back, p.Submit))
	p.Root = tui.NewVBox(view.NewContent(window), p.status)

	return p
}
