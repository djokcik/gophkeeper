package savepage

import (
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/view"
)

type LoginPasswordSagePage struct {
	view.PageHooks
	Root *tui.Box

	loginField    *tui.Entry
	passwordField *tui.Entry

	Submit *tui.Button
	Back   *tui.Button
}

func (p LoginPasswordSagePage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.loginField, p.passwordField, p.Submit, p.Back}
}

func (p LoginPasswordSagePage) GetRoot() tui.Widget {
	return p.Root
}

func (p LoginPasswordSagePage) OnActivated(fn func(b *tui.Button)) {
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

func NewLoginPasswordSagePage() *LoginPasswordSagePage {
	p := &LoginPasswordSagePage{Back: view.NewBackButton()}

	p.loginField = tui.NewEntry()
	p.loginField.SetFocused(true)
	formKey := tui.NewGrid(0, 0)
	formKey.AppendRow(p.loginField)
	loginBlock := tui.NewHBox(formKey)
	loginBlock.SetTitle("Логин")
	loginBlock.SetBorder(true)

	p.passwordField = tui.NewEntry()
	p.passwordField.SetEchoMode(tui.EchoModePassword)
	formKey = tui.NewGrid(0, 0)
	formKey.AppendRow(p.passwordField)
	passwordBlock := tui.NewHBox(formKey)
	passwordBlock.SetTitle("Пароль")
	passwordBlock.SetBorder(true)

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
