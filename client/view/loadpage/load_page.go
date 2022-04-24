package loadpage

import (
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/view"
)

type LoadPage struct {
	view.PageHooks
	Root *tui.Box

	keyField  *tui.Entry
	listField *tui.List

	Submit *tui.Button
	Back   *tui.Button
}

func (p LoadPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.Submit, p.Back}
}

func (p LoadPage) GetRoot() tui.Widget {
	return p.Root
}

func (p LoadPage) OnActivated(fn func(b *tui.Button)) {
	for _, button := range []*tui.Button{p.Back, p.Submit} {
		if button == p.Submit {
			if p.keyField.Text() == "" {
				return
			}

			continue
		}

		button.OnActivated(func(b *tui.Button) { fn(b) })
	}
}

func NewLoadPage() *LoadPage {
	p := &LoadPage{}

	p.keyField = tui.NewEntry()
	p.keyField.SetFocused(true)
	formKey := tui.NewGrid(0, 0)
	formKey.AppendRow(p.keyField)
	keyBlock := tui.NewHBox(formKey)
	keyBlock.SetTitle("Ключ")
	keyBlock.SetBorder(true)

	p.listField = tui.NewList()
	p.listField.AddItems("Пара логин/пароль")
	p.listField.AddItems("Текстовые данные")
	p.listField.AddItems("Бинарные данные")
	p.listField.AddItems("Банковская карта")
	p.listField.Select(0)
	p.listField.SetFocused(true)
	typeBlock := tui.NewVBox(p.listField)
	typeBlock.SetTitle("Тип")
	typeBlock.SetBorder(true)

	submit := tui.NewButton("[Получить]")
	p.Submit = submit
	p.Back = view.NewBackButton()

	buttons := tui.NewHBox(
		tui.NewPadder(1, 0, p.Back),
		tui.NewSpacer(),
		tui.NewPadder(1, 0, p.Submit),
	)

	window := tui.NewVBox(
		tui.NewPadder(10, 0, tui.NewLabel(view.Logo)),
		tui.NewSpacer(),
		tui.NewLabel("Выберите тип и ключ"),
		keyBlock,
		tui.NewSpacer(),
		typeBlock,
		tui.NewLabel(""),
		buttons,
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
