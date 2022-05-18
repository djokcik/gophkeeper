package loadpage

import (
	"context"
	"fmt"
	"github.com/djokcik/gophkeeper/client/registry"
	"github.com/djokcik/gophkeeper/client/view"
	"github.com/djokcik/gophkeeper/pkg/logging"
	"github.com/marcusolsson/tui-go"
)

// BankCardLoadPage is widget for Load BankCard
type BankCardLoadPage struct {
	view.PageHooks

	serviceRegistry registry.ClientServiceRegistry
	Root            *tui.Box

	keyField *tui.Entry
	result   *tui.Label
	status   *tui.StatusBar

	Submit *tui.Button
	Back   *tui.Button
}

// GetFocusChain returns list of focused widgets
func (p BankCardLoadPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.Submit, p.Back}
}

// GetRoot return Root winget element
func (p BankCardLoadPage) GetRoot() tui.Widget {
	return p.Root
}

// OnActivated call one time. Needed for navigate between pages
func (p BankCardLoadPage) OnActivated(fn func(b *tui.Button)) {
	service := p.serviceRegistry.GetRecordBankCardService()

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
			log.Error().Err(err).Msg("Submit: invalid LoadRecordByKey")

			return
		}

		p.status.SetText("Банковская карта успешно получена")
		p.result.SetText(fmt.Sprintf(
			"Номер карты: %s\nГод: %s\nCVV: %s\nПримечания: %s\n",
			data.CardNumber,
			data.Year,
			data.CVV,
			data.Comment,
		))
	})
}

func NewBankCardLoadPage(serviceRegistry registry.ClientServiceRegistry) *BankCardLoadPage {
	p := &BankCardLoadPage{
		serviceRegistry: serviceRegistry,
		Back:            view.NewBackButton(),
		Submit:          view.NewLoadButton(),
		status:          view.NewStatusLabel(),
		result:          view.NewResultLabel(),
	}

	window := view.NewWindowBlock("Найти банковскую карту")

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
