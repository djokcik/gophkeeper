package view

import (
	"github.com/marcusolsson/tui-go"
)

type SagePage struct {
	Root *tui.Box

	box *tui.Box

	LoginPassword *tui.Button
	TextButton    *tui.Button
	BinButton     *tui.Button
	CardButton    *tui.Button

	Back *tui.Button
}

func (p SagePage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.LoginPassword, p.TextButton, p.BinButton, p.CardButton, p.Back}
}

func (p SagePage) GetRoot() *tui.Box {
	return p.Root
}

func (p SagePage) GetButtons() []*tui.Button {
	return []*tui.Button{p.LoginPassword, p.TextButton, p.BinButton, p.CardButton, p.Back}
}

func (p SagePage) OnActivated(fn func(b *tui.Button)) {
	for _, button := range p.GetButtons() {
		button.OnActivated(func(b *tui.Button) { fn(b) })
	}
}

func NewSagePage() *SagePage {
	loginPassword := tui.NewButton("[Пара логин/пароль]")
	textButton := tui.NewButton("[Текстовые данные]")
	binButton := tui.NewButton("[Бинарные данные]")
	cardButton := tui.NewButton("[Банковская карта]")
	back := tui.NewButton("[Назад]")

	var rows = []tui.Widget{
		loginPassword,
		textButton,
		binButton,
		cardButton,
		//back,
	}

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, back),
	)

	table := tui.NewTable(0, 0)
	table.SetFocused(true)
	OnActivateButtonInRowTable(table, rows)
	for _, row := range rows {
		table.AppendRow(row)
	}

	box := tui.NewVBox(loginPassword,
		tui.NewSpacer(),
		textButton,
		tui.NewSpacer(),
		binButton,
		tui.NewSpacer(),
		cardButton)
	box.SetFocused(true)

	window := tui.NewVBox(
		tui.NewPadder(10, 0, tui.NewLabel(Logo)),
		tui.NewSpacer(),
		tui.NewLabel("Какой формат данных хотите сохранить? (Используйте TAB для навигации)"),
		tui.NewLabel(""),
		box,
		//table,
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

	return &SagePage{
		Root: root,

		box: box,

		CardButton:    cardButton,
		BinButton:     binButton,
		TextButton:    textButton,
		LoginPassword: loginPassword,
		Back:          back,
	}
}
