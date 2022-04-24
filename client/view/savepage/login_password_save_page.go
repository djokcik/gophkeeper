package savepage

import (
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/view"
)

type LoginPasswordSavePage struct {
	view.PageHooks
	Root *tui.Box

	loginField    *tui.Entry
	passwordField *tui.Entry

	Submit *tui.Button
	Back   *tui.Button
}

func (p LoginPasswordSavePage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.loginField, p.passwordField, p.Submit, p.Back}
}

func (p LoginPasswordSavePage) GetRoot() tui.Widget {
	return p.Root
}

func (p LoginPasswordSavePage) OnActivated(fn func(b *tui.Button)) {
	for _, button := range []*tui.Button{p.Back, p.Submit} {
		if button == p.Submit {
			if p.loginField.Text() == "" || p.passwordField.Text() == "" {
				return
			}

			continue
		}

		button.OnActivated(func(b *tui.Button) { fn(b) })
	}
}

func NewLoginPasswordSagePage() *LoginPasswordSavePage {
	p := &LoginPasswordSavePage{Back: view.NewBackButton()}

	loginField, loginBlock := view.NewEditBlock("Логин")
	p.loginField = loginField
	p.loginField.SetFocused(true)

	passwordField, passwordBlock := view.NewEditBlock("Пароль")
	p.passwordField = passwordField
	p.passwordField.SetEchoMode(tui.EchoModePassword)

	submit := tui.NewButton("[Сохранить]")
	p.Submit = submit
	p.Back = view.NewBackButton()

	buttons := tui.NewHBox(
		tui.NewPadder(1, 0, p.Back),
		tui.NewSpacer(),
		tui.NewPadder(1, 0, p.Submit),
	)

	window := tui.NewVBox(
		tui.NewPadder(10, 0, tui.NewLabel(view.Logo)),
		tui.NewSpacer(),
		tui.NewLabel("Сохраните логин и пароль"),
		loginBlock,
		tui.NewSpacer(),
		passwordBlock,
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

	p.Root = tui.NewVBox(content)

	return p
}
