package loadpage

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/view"
)

type BankCardLoadPage struct {
	view.PageHooks
	Root *tui.Box

	keyField *tui.Entry
	result   *tui.Label

	Submit *tui.Button
	Back   *tui.Button
}

func (p BankCardLoadPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.Submit, p.Back}
}

func (p BankCardLoadPage) GetRoot() tui.Widget {
	return p.Root
}

func (p BankCardLoadPage) OnActivated(fn func(b *tui.Button)) {
	for _, button := range []*tui.Button{p.Submit, p.Back} {
		button.OnActivated(func(b *tui.Button) {
			if b == p.Submit {
				if p.keyField.Text() == "" {
					return
				}

				if p.keyField.Text() == "test" {
					p.result.SetText(fmt.Sprintf("Не удалось найти карту по ключу - %s", p.keyField.Text()))

					return
				}

				p.result.SetText(fmt.Sprintf("Номер карты: 3000 5111 4123 4412\nГод: 2027\ncvv: 413"))

				return
			}

			fn(b)
		})
	}
}

func NewBankCardLoadPage() *BankCardLoadPage {
	p := &BankCardLoadPage{Back: view.NewBackButton()}

	keyField, keyBlock := view.NewEditBlock("Ключ")
	p.keyField = keyField
	p.keyField.SetFocused(true)

	p.Submit = tui.NewButton("[Найти]")
	p.Back = view.NewBackButton()
	p.result = tui.NewLabel("")

	box := tui.NewVBox(tui.NewLabel("Результат:\n"), p.result)
	box.SetBorder(true)

	window := tui.NewVBox(
		tui.NewPadder(10, 0, tui.NewLabel(view.Logo)),
		tui.NewSpacer(),
		tui.NewLabel("\nНайти бинарные данные\n"),
		keyBlock,
		tui.NewLabel(""),
		tui.NewHBox(
			tui.NewSpacer(),
			tui.NewPadder(1, 0, p.Submit),
		),
		box,
		tui.NewHBox(
			tui.NewPadder(1, 0, p.Back),
			tui.NewSpacer(),
		),
	)
	window.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)
	content := tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())

	p.Root = tui.NewVBox(content)

	return p
}
