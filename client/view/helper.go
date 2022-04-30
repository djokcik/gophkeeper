package view

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
)

func NewBackButton() *tui.Button {
	return tui.NewButton("[Назад]")
}

func NewLoadButton() *tui.Button {
	return tui.NewButton("[Найти]")
}

func NewResultLabel() *tui.Label {
	return tui.NewLabel("")
}

func NewStatusLabel() *tui.StatusBar {
	return tui.NewStatusBar("")
}

func NewSaveButton() *tui.Button {
	return tui.NewButton("[Сохранить]")
}

func NewContent(window *tui.Box) *tui.Box {
	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)

	return tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())
}

func NewButtons(back *tui.Button, submit *tui.Button) *tui.Box {
	box := tui.NewHBox()

	if back != nil {
		box.Append(tui.NewPadder(1, 0, back))
	}

	box.Append(tui.NewSpacer())

	if submit != nil {
		box.Append(tui.NewPadder(1, 0, submit))
	}

	return box
}

func NewWindowBlock(label string) *tui.Box {
	block := tui.NewVBox(
		tui.NewPadder(10, 0, tui.NewLabel(Logo)),
		tui.NewSpacer(),
		tui.NewLabel(fmt.Sprintf("\n%s\n", label)),
	)

	block.SetBorder(true)

	return block
}

func NewEditBlockWithWindow(title string, window *tui.Box) *tui.Entry {
	entry, entryBlock := NewEditBlock(title)
	window.Append(entryBlock)
	window.Append(tui.NewSpacer())

	return entry
}

func NewEditBlock(title string) (*tui.Entry, *tui.Box) {
	field := tui.NewEntry()
	formKey := tui.NewGrid(0, 0)
	formKey.AppendRow(field)
	block := tui.NewHBox(formKey)
	block.SetTitle(title)
	block.SetBorder(true)

	return field, block
}
