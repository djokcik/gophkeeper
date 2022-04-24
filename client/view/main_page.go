package view

import (
	"github.com/marcusolsson/tui-go"
)

type MainPage struct {
	Root *tui.Box

	SaveData *tui.Button
	LoadData *tui.Button
	Table    *tui.Table
}

func (p MainPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{}
}

func (p MainPage) GetRoot() *tui.Box {
	return p.Root
}

func (p MainPage) GetButtons() []*tui.Button {
	return []*tui.Button{p.SaveData, p.LoadData}
}

func (p MainPage) OnActivated(fn func(*tui.Button)) {
	for _, button := range p.GetButtons() {
		button.OnActivated(func(b *tui.Button) { fn(b) })
	}
}

func NewMainPage() *MainPage {
	saveData := tui.NewButton("[Сохранить данные]")
	loadData := tui.NewButton("[Получить данные]")

	var rows = []tui.Widget{saveData, loadData}

	table := tui.NewTable(0, 0)
	table.SetFocused(true)
	for _, row := range rows {
		table.AppendRow(row)
	}

	OnActivateButtonInRowTable(table, rows)

	window := tui.NewVBox(
		tui.NewPadder(10, 0, tui.NewLabel(Logo)),
		tui.NewSpacer(),
		tui.NewLabel("\nДобрый день, Иван."),
		tui.NewLabel(""),
		table,
	)
	window.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)
	content := tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())

	root := tui.NewVBox(content)

	return &MainPage{
		Root: root,

		SaveData: saveData,
		LoadData: loadData,
	}
}
