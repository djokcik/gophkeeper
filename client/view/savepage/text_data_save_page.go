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

// TextDataSavePage is widget for Save TextData
type TextDataSavePage struct {
	view.PageHooks

	serviceRegistry registry.ClientServiceRegistry
	Root            *tui.Box

	keyField      *tui.Entry
	textDataField *tui.Entry
	commentField  *tui.Entry
	status        *tui.StatusBar

	Submit *tui.Button
	Back   *tui.Button
}

// GetFocusChain returns list of focused widgets
func (p TextDataSavePage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.textDataField, p.commentField, p.Submit, p.Back}
}

// GetRoot return Root winget element
func (p TextDataSavePage) GetRoot() tui.Widget {
	return p.Root
}

// OnActivated call one time. Needed for navigate between pages
func (p TextDataSavePage) OnActivated(fn func(b *tui.Button)) {
	service := p.serviceRegistry.GetRecordTextDataService()

	p.Back.OnActivated(fn)
	p.Submit.OnActivated(func(b *tui.Button) {
		ctx, log := logging.GetCtxFileLogger(context.Background())

		if p.keyField.Text() == "" || p.textDataField.Text() == "" {
			p.status.SetText("Не все поля заполнены")
			return
		}

		data := clientmodels.RecordTextData{
			Text:    p.textDataField.Text(),
			Comment: p.commentField.Text(),
		}

		err := service.SaveRecord(ctx, p.keyField.Text(), data)
		if err != nil {
			p.status.SetText(fmt.Sprintf("Error: %s", err.Error()))
			log.Error().Err(err).Msg("Submit: invalid SaveRecord")

			return
		}

		p.status.SetText("Данные успешно сохранены")
	})
}

func NewTextDataSavePage(serviceRegistry registry.ClientServiceRegistry) *TextDataSavePage {
	p := &TextDataSavePage{
		Back:            view.NewBackButton(),
		Submit:          view.NewSaveButton(),
		status:          view.NewStatusLabel(),
		serviceRegistry: serviceRegistry,
	}

	window := view.NewWindowBlock("Сохранить текст по ключу")

	p.keyField = view.NewEditBlockWithWindow("Ключ", window)
	p.keyField.SetFocused(true)

	p.textDataField = view.NewEditBlockWithWindow("Текст", window)
	p.commentField = view.NewEditBlockWithWindow("Примечание", window)

	window.Append(view.NewButtons(p.Back, p.Submit))
	p.Root = tui.NewVBox(view.NewContent(window), p.status)

	return p
}
