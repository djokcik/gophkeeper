package view

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
)

// NewBackButton returns new back button
func NewBackButton() *tui.Button {
	return tui.NewButton("[Назад]")
}

// NewLoadButton returns new load button
func NewLoadButton() *tui.Button {
	return tui.NewButton("[Найти]")
}

// NewResultLabel returns new result label
func NewResultLabel() *tui.Label {
	return tui.NewLabel("")
}

// NewStatusLabel returns new status label
func NewStatusLabel() *tui.StatusBar {
	return tui.NewStatusBar("")
}

// NewSaveButton returns new save button
func NewSaveButton() *tui.Button {
	return tui.NewButton("[Сохранить]")
}

// NewContent returns new content box
func NewContent(window *tui.Box) *tui.Box {
	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)

	return tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())
}

// NewButtons returns new buttons box
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

// NewWindowBlock returns new window block box
func NewWindowBlock(label string) *tui.Box {
	return NewWindowBlockLabel(tui.NewLabel(fmt.Sprintf("\n%s\n", label)))
}

// NewWindowBlockLabel returns new window block label box
func NewWindowBlockLabel(label *tui.Label) *tui.Box {
	block := tui.NewVBox(
		tui.NewLabel(fmt.Sprintf("Дата сборки: %s. Версия: %s", BuildDate, BuildVersion)),
		tui.NewPadder(10, 0, tui.NewLabel(Logo)),
		tui.NewSpacer(),
		label,
	)

	block.SetBorder(true)

	return block
}

// NewEditBlockWithWindow returns new window and block entry
func NewEditBlockWithWindow(title string, window *tui.Box) *tui.Entry {
	entry, entryBlock := NewEditBlock(title)
	window.Append(entryBlock)
	window.Append(tui.NewSpacer())

	return entry
}

// NewEditBlock returns new edit block
func NewEditBlock(title string) (*tui.Entry, *tui.Box) {
	field := tui.NewEntry()
	formKey := tui.NewGrid(0, 0)
	formKey.AppendRow(field)
	block := tui.NewHBox(formKey)
	block.SetTitle(title)
	block.SetBorder(true)

	return field, block
}
