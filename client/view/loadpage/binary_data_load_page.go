package loadpage

import (
	"context"
	"fmt"
	"github.com/djokcik/gophkeeper/client/registry"
	"github.com/djokcik/gophkeeper/client/view"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/marcusolsson/tui-go"
	"os"
)

type BinaryDataLoadPage struct {
	view.PageHooks

	serviceRegistry registry.ClientServiceRegistry
	Root            *tui.Box

	keyField        *tui.Entry
	pathToFileField *tui.Entry
	result          *tui.Label
	status          *tui.StatusBar

	Submit *tui.Button
	Back   *tui.Button
}

func (p BinaryDataLoadPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.pathToFileField, p.Submit, p.Back}
}

func (p BinaryDataLoadPage) GetRoot() tui.Widget {
	return p.Root
}

func (p BinaryDataLoadPage) OnActivated(fn func(b *tui.Button)) {
	service := p.serviceRegistry.GetRecordBinaryDataService()

	p.Back.OnActivated(fn)
	p.Submit.OnActivated(func(b *tui.Button) {
		p.status.SetText("Загрузка...")
		ctx, log := logging.GetCtxFileLogger(context.Background())

		if p.keyField.Text() == "" || p.pathToFileField.Text() == "" {
			p.status.SetText("Не введен ключ")
			return
		}

		data, err := service.LoadRecordByKey(ctx, p.keyField.Text())
		if err != nil {
			p.status.SetText(fmt.Sprintf("Error: %s", err.Error()))
			log.Error().Err(err).Msg("Submit: invalid LoadRecordByKey")

			return
		}

		filepath := p.pathToFileField.Text()

		err = os.WriteFile(filepath, data.Data, 0777)
		if err != nil {
			p.status.SetText(fmt.Sprintf("Error: %s", err.Error()))
			log.Error().Err(err).Msg("Submit: invalid WriteFile")

			return
		}

		p.status.SetText("Запись успешно получена")
		p.result.SetText(fmt.Sprintf(
			"Данные сохранены в файл: %s\nПримечания: %s\n",
			filepath,
			data.Comment,
		))
	})
}

func NewBinaryDataLoadPage(serviceRegistry registry.ClientServiceRegistry) *BinaryDataLoadPage {
	p := &BinaryDataLoadPage{
		serviceRegistry: serviceRegistry,
		Back:            view.NewBackButton(),
		Submit:          view.NewLoadButton(),
		status:          view.NewStatusLabel(),
		result:          view.NewResultLabel(),
	}

	window := view.NewWindowBlock("Найти файл")

	p.keyField = view.NewEditBlockWithWindow("Ключ", window)
	p.keyField.SetFocused(true)

	p.pathToFileField = view.NewEditBlockWithWindow("Путь до файлу куда сохранить результат", window)
	p.pathToFileField.SetText("/tmp/temp.txt")

	box := tui.NewVBox(tui.NewLabel("Результат:\n"), p.result)
	box.SetBorder(true)

	window.Append(view.NewButtons(nil, p.Submit))
	window.Append(box)
	window.Append(view.NewButtons(p.Back, nil))

	p.Root = tui.NewVBox(view.NewContent(window), p.status)

	return p
}
