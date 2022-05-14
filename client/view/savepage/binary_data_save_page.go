package savepage

import (
	"context"
	"fmt"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/clientmodels"
	"gophkeeper/client/registry"
	"gophkeeper/client/view"
	"gophkeeper/pkg/logging"
	"os"
)

type BinaryDataSavePage struct {
	view.PageHooks

	serviceRegistry registry.ClientServiceRegistry
	Root            *tui.Box

	keyField        *tui.Entry
	pathToFileField *tui.Entry
	commentField    *tui.Entry
	status          *tui.StatusBar

	Submit *tui.Button
	Back   *tui.Button
}

func (p BinaryDataSavePage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.pathToFileField, p.commentField, p.Submit, p.Back}
}

func (p BinaryDataSavePage) GetRoot() tui.Widget {
	return p.Root
}

func (p BinaryDataSavePage) OnActivated(fn func(b *tui.Button)) {
	service := p.serviceRegistry.GetRecordBinaryDataService()

	p.Back.OnActivated(fn)
	p.Submit.OnActivated(func(b *tui.Button) {
		ctx, log := logging.GetCtxFileLogger(context.Background())

		if p.keyField.Text() == "" || p.pathToFileField.Text() == "" {
			p.status.SetText("Не все поля заполнены")
			return
		}

		bytes, err := os.ReadFile(p.pathToFileField.Text())
		if err != nil {
			p.status.SetText(fmt.Sprintf("Error ReadFile: %s", err.Error()))
			log.Error().Err(err).Msg("Submit: invalid ReadFile")

			return
		}

		data := clientmodels.RecordBinaryData{
			Data:    bytes,
			Comment: p.commentField.Text(),
		}

		err = service.SaveRecord(ctx, p.keyField.Text(), data)
		if err != nil {
			p.status.SetText(fmt.Sprintf("Error: %s", err.Error()))
			log.Error().Err(err).Msg("Submit: invalid SaveRecord")

			return
		}

		p.status.SetText("Данные успешно сохранены")
	})
}

func NewBinaryDataSavePage(serviceRegistry registry.ClientServiceRegistry) *BinaryDataSavePage {
	p := &BinaryDataSavePage{
		Back:            view.NewBackButton(),
		Submit:          view.NewSaveButton(),
		status:          view.NewStatusLabel(),
		serviceRegistry: serviceRegistry,
	}

	window := view.NewWindowBlock("Сохранить текст по ключу")

	p.keyField = view.NewEditBlockWithWindow("Ключ", window)
	p.keyField.SetFocused(true)

	p.pathToFileField = view.NewEditBlockWithWindow("Путь до файла", window)
	p.commentField = view.NewEditBlockWithWindow("Примечание", window)

	window.Append(view.NewButtons(p.Back, p.Submit))
	p.Root = tui.NewVBox(view.NewContent(window), p.status)

	return p
}
