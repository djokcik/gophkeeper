package view

import "github.com/marcusolsson/tui-go"

type LoginRegisterPage struct {
	Root *tui.Box

	user     *tui.Entry
	password *tui.Entry
	Login    *tui.Button
	Register *tui.Button
}

func (p LoginRegisterPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.user, p.password, p.Login, p.Register}
}

func (p LoginRegisterPage) GetRoot() *tui.Box {
	return p.Root
}

func (p LoginRegisterPage) GetButtons() []*tui.Button {
	return []*tui.Button{p.Login, p.Register}
}

func (p LoginRegisterPage) OnActivated(fn func(b *tui.Button)) {
	p.Login.OnActivated(func(b *tui.Button) {
		if p.user.Text() == "" || p.password.Text() == "" {
			return
		}

		fn(b)
	})

	p.Register.OnActivated(func(b *tui.Button) {
		if p.user.Text() == "" || p.password.Text() == "" {
			return
		}

		fn(b)
	})
}

func NewLoginRegisterPage() *LoginRegisterPage {
	user := tui.NewEntry()
	user.SetFocused(true)
	formUser := tui.NewGrid(0, 0)
	formUser.AppendRow(tui.NewLabel("Логин"))
	formUser.AppendRow(user)
	formUserBox := tui.NewHBox(formUser)
	formUserBox.SetBorder(true)

	password := tui.NewEntry()
	password.SetEchoMode(tui.EchoModePassword)
	formPassword := tui.NewGrid(0, 0)
	formPassword.AppendRow(tui.NewLabel("Пароль"))
	formPassword.AppendRow(password)
	formPasswordBox := tui.NewHBox(formPassword)
	formPasswordBox.SetBorder(true)

	form := tui.NewGrid(0, 0)
	form.AppendRow(formUserBox, formPasswordBox)

	login := tui.NewButton("[Вход]")
	register := tui.NewButton("[Регистрация]")

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, login),
		tui.NewPadder(1, 0, register),
	)

	window := tui.NewVBox(
		tui.NewPadder(10, 0, tui.NewLabel(Logo)),
		tui.NewPadder(18, 1, tui.NewLabel("Добро пожаловать в GophKeeper!")),
		tui.NewPadder(1, 1, form),
		buttons,
	)
	window.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)
	content := tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())

	root := tui.NewVBox(content)

	return &LoginRegisterPage{
		Root: root,

		Register: register,
		Login:    login,
		password: password,
		user:     user,
	}
}
