package loadpage

import (
	"context"
	"fmt"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/registry"
	"gophkeeper/client/view"
	"gophkeeper/pkg/logging"
)

type TextDataLoadPage struct {
	view.PageHooks

	serviceRegistry registry.ClientServiceRegistry
	Root            *tui.Box

	keyField *tui.Entry
	result   *tui.Label
	status   *tui.StatusBar

	Submit *tui.Button
	Back   *tui.Button
}

func (p TextDataLoadPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.Submit, p.Back}
}

func (p TextDataLoadPage) GetRoot() tui.Widget {
	return p.Root
}

func (p TextDataLoadPage) OnActivated(fn func(b *tui.Button)) {
	service := p.serviceRegistry.GetRecordTextDataService()

	p.Back.OnActivated(fn)
	p.Submit.OnActivated(func(b *tui.Button) {
		p.status.SetText("Загрузка...")
		ctx, log := logging.GetCtxFileLogger(context.Background())

		if p.keyField.Text() == "" {
			p.status.SetText("Не введен ключ")
			return
		}

		data, err := service.LoadRecordByKey(ctx, p.keyField.Text())
		if err != nil {
			p.status.SetText(fmt.Sprintf("Error: %s", err.Error()))
			log.Error().Err(err).Msg("Submit: invalid LoadRecordPersonalDataByKey")

			return
		}

		p.status.SetText("Запись успешно получена")
		p.result.SetText(fmt.Sprintf(
			"Текст: %s\nПримечания: %s\n",
			data.Text,
			data.Comment,
		))
	})
}

func NewTextDataLoadPage(serviceRegistry registry.ClientServiceRegistry) *TextDataLoadPage {
	p := &TextDataLoadPage{
		serviceRegistry: serviceRegistry,
		Back:            view.NewBackButton(),
		Submit:          view.NewLoadButton(),
		status:          view.NewStatusLabel(),
		result:          view.NewResultLabel(),
	}

	window := view.NewWindowBlock("Найти запись по ключу")

	p.keyField = view.NewEditBlockWithWindow("Ключ", window)
	p.keyField.SetFocused(true)

	box := tui.NewVBox(tui.NewLabel("Результат:\n"), p.result)
	box.SetBorder(true)

	window.Append(view.NewButtons(nil, p.Submit))
	window.Append(box)
	window.Append(view.NewButtons(p.Back, nil))

	p.Root = tui.NewVBox(view.NewContent(window), p.status)

	return p
}
