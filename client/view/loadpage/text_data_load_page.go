package loadpage

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/view"
)

type TextDataLoadPage struct {
	view.PageHooks
	Root *tui.Box

	keyField *tui.Entry
	result   *tui.Label

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
	for _, button := range []*tui.Button{p.Submit, p.Back} {
		button.OnActivated(func(b *tui.Button) {
			if b == p.Submit {
				if p.keyField.Text() == "" {
					return
				}

				if p.keyField.Text() == "test" {
					p.result.SetText(fmt.Sprintf("Не удалось найти данные по ключу - %s", p.keyField.Text()))

					return
				}

				p.result.SetText(fmt.Sprintf("Данные: %s", p.keyField.Text()))

				return
			}

			fn(b)
		})
	}
}

func NewTextDataLoadPage() *TextDataLoadPage {
	p := &TextDataLoadPage{Back: view.NewBackButton()}

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
		tui.NewLabel("\nПоказать данные по ключу\n"),
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
