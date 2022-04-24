package view

import "github.com/marcusolsson/tui-go"

func NewBackButton() *tui.Button {
	return tui.NewButton("[Назад]")
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
