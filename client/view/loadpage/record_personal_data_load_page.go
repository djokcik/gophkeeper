package loadpage

import (
	"context"
	"fmt"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/registry"
	"gophkeeper/client/view"
	"gophkeeper/pkg/logging"
)

type RecordPersonalDataLoadPage struct {
	view.PageHooks
	serviceRegistry registry.ClientServiceRegistry

	Root *tui.Box

	keyField *tui.Entry
	result   *tui.Label
	status   *tui.StatusBar

	Submit *tui.Button
	Back   *tui.Button
}

func (p RecordPersonalDataLoadPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.Submit, p.Back}
}

func (p RecordPersonalDataLoadPage) GetRoot() tui.Widget {
	return p.Root
}

func (p RecordPersonalDataLoadPage) OnActivated(fn func(b *tui.Button)) {
	service := p.serviceRegistry.GetRecordPersonalDataService()

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
			"Логин: %s\nПароль: %s\nURL-адрес: %s\nПримечания: %s\n",
			data.Username,
			data.Password,
			data.URL,
			data.Comment,
		))

		return
	})
}

func NewRecordPersonalDataLoadPage(serviceRegistry registry.ClientServiceRegistry) *RecordPersonalDataLoadPage {
	p := &RecordPersonalDataLoadPage{
		serviceRegistry: serviceRegistry,
		Back:            view.NewBackButton(),
		Submit:          view.NewLoadButton(),
		status:          view.NewStatusLabel(),
		result:          view.NewResultLabel(),
	}

	wBlock := view.NewWindowBlock("Найти запись по ключу")

	p.keyField = view.NewEditBlockWithWindow("Ключ", wBlock)
	p.keyField.SetFocused(true)

	box := tui.NewVBox(tui.NewLabel("Результат:\n"), p.result)
	box.SetBorder(true)

	wBlock.Append(view.NewButtons(nil, p.Submit))
	wBlock.Append(box)
	wBlock.Append(view.NewButtons(p.Back, nil))

	p.Root = tui.NewVBox(view.NewContent(wBlock), p.status)

	return p
}
