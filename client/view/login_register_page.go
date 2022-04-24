package view

import (
	"context"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/registry"
)

type LoginRegisterPage struct {
	PageHooks
	serviceRegistry registry.ServiceRegistry

	Root *tui.Box

	user     *tui.Entry
	password *tui.Entry
	status   *tui.StatusBar

	Login    *tui.Button
	Register *tui.Button
}

func (p LoginRegisterPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.user, p.password, p.Login, p.Register}
}

func (p LoginRegisterPage) GetRoot() tui.Widget {
	return p.Root
}

func (p LoginRegisterPage) OnActivated(fn func(b *tui.Button)) {
	loginService := p.serviceRegistry.GetAuthService()
	userService := p.serviceRegistry.GetUserService()

	p.Login.OnActivated(func(b *tui.Button) {
		ctx := context.Background()

		if p.user.Text() == "" || p.password.Text() == "" {
			return
		}

		p.status.SetText("Загрузка...")
		user, err := loginService.Login(ctx, p.user.Text(), p.password.Text())
		if err != nil {
			p.status.SetText(err.Error())
			return
		}

		userService.SaveUser(user)

		fn(b)
	})

	p.Register.OnActivated(func(b *tui.Button) {
		ctx := context.Background()

		if p.user.Text() == "" || p.password.Text() == "" {
			return
		}

		p.status.SetText("Загрузка...")
		user, err := loginService.Register(ctx, p.user.Text(), p.password.Text())
		if err != nil {
			p.status.SetText(err.Error())
			return
		}

		userService.SaveUser(user)

		fn(b)
	})
}

func NewLoginRegisterPage(serviceRegistry registry.ServiceRegistry) *LoginRegisterPage {
	page := &LoginRegisterPage{serviceRegistry: serviceRegistry}

	user := tui.NewEntry()
	user.SetFocused(true)
	formUser := tui.NewGrid(0, 0)
	formUser.AppendRow(tui.NewLabel("Логин"))
	formUser.AppendRow(user)
	formUserBox := tui.NewHBox(formUser)
	formUserBox.SetBorder(true)
	page.user = user

	password := tui.NewEntry()
	password.SetEchoMode(tui.EchoModePassword)
	formPassword := tui.NewGrid(0, 0)
	formPassword.AppendRow(tui.NewLabel("Пароль"))
	formPassword.AppendRow(password)
	formPasswordBox := tui.NewHBox(formPassword)
	formPasswordBox.SetBorder(true)
	page.password = password

	form := tui.NewGrid(0, 0)
	form.AppendRow(formUserBox, formPasswordBox)

	page.Login = tui.NewButton("[Вход]")
	page.Register = tui.NewButton("[Регистрация]")

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, page.Login),
		tui.NewPadder(1, 0, page.Register),
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

	page.status = tui.NewStatusBar("Введите логин и пароль")
	page.Root = tui.NewVBox(content, page.status)

	return page
}
