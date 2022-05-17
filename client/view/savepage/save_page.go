package savepage

import (
	"github.com/djokcik/gophkeeper/client/view"
	"github.com/marcusolsson/tui-go"
)

type SavePage struct {
	view.PageHooks
	Root *tui.Box

	LoginPassword *tui.Button
	TextButton    *tui.Button
	BinButton     *tui.Button
	CardButton    *tui.Button

	Back *tui.Button
}

func (p SavePage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.LoginPassword, p.TextButton, p.BinButton, p.CardButton, p.Back}
}

func (p SavePage) GetRoot() tui.Widget {
	return p.Root
}

func (p SavePage) OnActivated(fn func(b *tui.Button)) {
	buttons := []*tui.Button{p.LoginPassword, p.TextButton, p.BinButton, p.CardButton, p.Back}

	for _, button := range buttons {
		button.OnActivated(func(b *tui.Button) { fn(b) })
	}
}

func NewSavePage() *SavePage {
	p := &SavePage{Back: view.NewBackButton()}

	window := view.NewWindowBlock("Выберите формат данных? (Используйте TAB для навигации)")

	p.LoginPassword = tui.NewButton("[Пара логин/пароль]")
	p.TextButton = tui.NewButton("[Текстовые данные]")
	p.BinButton = tui.NewButton("[Бинарные данные]")
	p.CardButton = tui.NewButton("[Банковская карта]")

	box := tui.NewVBox(
		p.LoginPassword,
		p.TextButton,
		p.BinButton,
		p.CardButton,
	)
	box.SetFocused(true)

	window.Append(box)
	window.Append(tui.NewLabel(""))
	window.Append(p.Back)
	p.Root = tui.NewVBox(view.NewContent(window))

	return p
}
