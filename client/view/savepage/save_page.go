package savepage

import (
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/view"
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

func NewSagePage() *SavePage {
	p := &SavePage{Back: view.NewBackButton()}

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

	window := tui.NewVBox(
		tui.NewPadder(10, 0, tui.NewLabel(view.Logo)),
		tui.NewSpacer(),
		tui.NewLabel("Какой формат данных хотите сохранить? (Используйте TAB для навигации)"),
		tui.NewLabel(""),
		box,
		tui.NewLabel(""),
		tui.NewHBox(
			tui.NewSpacer(),
			tui.NewPadder(1, 0, p.Back),
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
