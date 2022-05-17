package removepage

import (
	"context"
	"fmt"
	"github.com/djokcik/gophkeeper/client/registry"
	"github.com/djokcik/gophkeeper/client/view"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/marcusolsson/tui-go"
)

var (
	PersonalDataItem = "Пары логин/пароль"
	TextDataItem     = "Произвольные текстовые данные"
	BinDataItem      = "Произвольные бинарные данные"
	CardDataItem     = "Данные банковских карт"
)

type RemovePage struct {
	view.PageHooks
	serviceRegistry registry.ClientServiceRegistry

	Root *tui.Box

	keyField *tui.Entry

	status *tui.StatusBar
	list   *tui.List

	Remove *tui.Button
	Back   *tui.Button
}

func (p RemovePage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.list, p.Remove, p.Back}
}

func (p RemovePage) GetRoot() tui.Widget {
	return p.Root
}

func (p RemovePage) OnActivated(fn func(b *tui.Button)) {
	p.Back.OnActivated(fn)
	p.Remove.OnActivated(func(b *tui.Button) {
		p.status.SetText("Загрузка...")
		ctx, log := logging.GetCtxFileLogger(context.Background())

		if p.keyField.Text() == "" {
			p.status.SetText("Не введен ключ")
			return
		}

		var err error

		switch p.list.SelectedItem() {
		case PersonalDataItem:
			err = p.serviceRegistry.GetRecordPersonalDataService().RemoveRecordByKey(ctx, p.keyField.Text())
		}

		if err != nil {
			p.status.SetText(fmt.Sprintf("Error: %s", err.Error()))
			log.Error().Err(err).Msg("Submit: invalid LoadRecordPersonalDataByKey")

			return
		}

		p.status.SetText("Запись успешно удалена")
	})
}

func NewRemovePage(serviceRegistry registry.ClientServiceRegistry) *RemovePage {
	p := &RemovePage{
		Back:            view.NewBackButton(),
		Remove:          tui.NewButton("[Удалить]"),
		status:          view.NewStatusLabel(),
		serviceRegistry: serviceRegistry,
	}

	window := view.NewWindowBlock("Удалить запись по ключу")

	p.keyField = view.NewEditBlockWithWindow("Ключ", window)
	p.keyField.SetFocused(true)

	box := tui.NewVBox()
	box.SetTitle("Тип операции")
	box.SetBorder(true)

	p.list = tui.NewList()
	p.list.AddItems(
		PersonalDataItem,
		TextDataItem,
		BinDataItem,
		CardDataItem,
	)
	p.list.Select(0)
	box.Append(p.list)

	window.Append(box)
	window.Append(view.NewButtons(p.Back, p.Remove))
	p.Root = tui.NewVBox(view.NewContent(window), p.status)

	return p
}
