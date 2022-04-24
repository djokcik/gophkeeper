package loadpage

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/view"
)

type LoginPasswordLoadPage struct {
	view.PageHooks
	Root *tui.Box

	loginField *tui.Entry
	result     *tui.Label

	Submit *tui.Button
	Back   *tui.Button
}

func (p LoginPasswordLoadPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.loginField, p.Submit, p.Back}
}

func (p LoginPasswordLoadPage) GetRoot() tui.Widget {
	return p.Root
}

func (p LoginPasswordLoadPage) OnActivated(fn func(b *tui.Button)) {
	for _, button := range []*tui.Button{p.Submit, p.Back} {
		button.OnActivated(func(b *tui.Button) {
			if b == p.Submit {
				if p.loginField.Text() == "" {
					return
				}

				if p.loginField.Text() == "test" {
					p.result.SetText(fmt.Sprintf("Не удалось найти пароль по логину - %s", p.loginField.Text()))

					return
				}

				p.result.SetText(fmt.Sprintf("Пароль: %s", p.loginField.Text()))

				return
			}

			fn(b)
		})
	}
}

func NewLoginPasswordLoadPage() *LoginPasswordLoadPage {
	p := &LoginPasswordLoadPage{Back: view.NewBackButton()}

	loginField, loginBlock := view.NewEditBlock("Логин")
	p.loginField = loginField
	p.loginField.SetFocused(true)

	p.Submit = tui.NewButton("[Найти]")
	p.Back = view.NewBackButton()
	p.result = tui.NewLabel("")

	box := tui.NewVBox(tui.NewLabel("Результат:\n"), p.result)
	box.SetBorder(true)

	window := tui.NewVBox(
		tui.NewPadder(10, 0, tui.NewLabel(view.Logo)),
		tui.NewSpacer(),
		tui.NewLabel("\nПоказать пароль по логину\n"),
		loginBlock,
		tui.NewLabel(""),
		tui.NewHBox(
			tui.NewSpacer(),
			tui.NewPadder(1, 0, p.Submit),
		),
		box,
		tui.NewHBox(
			tui.NewPadder(1, 0, p.Back),
			tui.NewSpacer(),
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
