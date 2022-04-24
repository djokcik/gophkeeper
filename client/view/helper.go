package view

import "github.com/marcusolsson/tui-go"

func OnActivateButtonInRowTable(table *tui.Table, rows []tui.Widget) {
	table.OnItemActivated(func(table *tui.Table) {
		row := rows[table.Selected()]
		if b, ok := row.(*tui.Button); ok {
			b.SetFocused(true)
			b.OnKeyEvent(tui.KeyEvent{Key: tui.KeyEnter})
			b.SetFocused(false)
		}
	})
}
