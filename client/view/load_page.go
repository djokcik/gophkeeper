package view

import (
	"github.com/marcusolsson/tui-go"
)

type LoadPage struct {
	Root *tui.Box

	keyField  *tui.Entry
	listField *tui.List

	KeyBlock  *tui.Box
	TypeBlock *tui.Box

	Back   *tui.Button
	Submit *tui.Button
}

func (p LoadPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.KeyBlock, p.Submit, p.Back}
}

func (p LoadPage) GetRoot() *tui.Box {
	return p.Root
}

func (p LoadPage) GetButtons() []*tui.Button {
	return []*tui.Button{p.Back, p.Submit}
}

func (p LoadPage) OnActivated(fn func(b *tui.Button)) {
	for _, button := range p.GetButtons() {
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
	key := tui.NewEntry()
	key.SetFocused(true)
	formKey := tui.NewGrid(0, 0)
	formKey.AppendRow(key)
	formKeyBox := tui.NewHBox(formKey)
	formKeyBox.SetTitle("Ключ")
	formKeyBox.SetBorder(true)

	list := tui.NewList()
	list.AddItems("Пара логин/пароль")
	list.AddItems("Текстовые данные")
	list.AddItems("Бинарные данные")
	list.AddItems("Банковская карта")
	list.Select(0)
	list.SetFocused(true)
	typeBlock := tui.NewVBox(list)
	typeBlock.SetTitle("Тип")
	typeBlock.SetBorder(true)

	back := tui.NewButton("[Назад]")
	submit := tui.NewButton("[Получить]")

	buttons := tui.NewHBox(
		tui.NewPadder(1, 0, back),
		tui.NewSpacer(),
		tui.NewPadder(1, 0, submit),
	)

	window := tui.NewVBox(
		tui.NewPadder(10, 0, tui.NewLabel(Logo)),
		tui.NewSpacer(),
		tui.NewLabel("Выберите тип и ключ"),
		formKeyBox,
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

	root := tui.NewVBox(content)

	return &LoadPage{
		Root: root,

		keyField:  key,
		listField: list,

		KeyBlock:  formKeyBox,
		TypeBlock: typeBlock,

		Back:   back,
		Submit: submit,
	}
}
