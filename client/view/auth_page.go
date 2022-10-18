package view

import (
	"context"
	"github.com/djokcik/gophkeeper/client/registry"
	"github.com/marcusolsson/tui-go"
)

// AuthPage is widget for authorization page
type AuthPage struct {
	PageHooks
	serviceRegistry registry.ClientServiceRegistry

	Root *tui.Box

	user     *tui.Entry
	password *tui.Entry
	status   *tui.StatusBar

	Login    *tui.Button
	Register *tui.Button
}

// GetFocusChain returns list of focused widgets
func (p AuthPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.user, p.password, p.Login, p.Register}
}

// GetRoot return Root winget element
func (p AuthPage) GetRoot() tui.Widget {
	return p.Root
}

// OnActivated call one time. Needed for navigate between pages
func (p AuthPage) OnActivated(fn func(b *tui.Button)) {
	loginService := p.serviceRegistry.GetAuthService()
	userService := p.serviceRegistry.GetUserService()

	p.Login.OnActivated(func(b *tui.Button) {
		ctx := context.Background()

		if p.user.Text() == "" || p.password.Text() == "" {
			return
		}

		p.status.SetText("Загрузка...")
		user, err := loginService.SignIn(ctx, p.user.Text(), p.password.Text())
		if err != nil {
			p.status.SetText(err.Error())
			return
		}

		err = userService.SaveUser(user)
		if err != nil {
			p.status.SetText(err.Error())
			return
		}

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

		err = userService.SaveUser(user)
		if err != nil {
			p.status.SetText(err.Error())
			return
		}

		fn(b)
	})
}

func NewLoginRegisterPage(serviceRegistry registry.ClientServiceRegistry) *AuthPage {
	p := &AuthPage{
		serviceRegistry: serviceRegistry,
		status:          tui.NewStatusBar("Введите логин и пароль"),
	}

	wBox := NewWindowBlock("")

	p.user = NewEditBlockWithWindow("Логин", wBox)
	p.user.SetFocused(true)

	p.password = NewEditBlockWithWindow("Пароль", wBox)
	p.password.SetEchoMode(tui.EchoModePassword)

	p.Login = tui.NewButton("[Вход]")
	p.Register = tui.NewButton("[Регистрация]")

	wBox.Append(tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, p.Login),
		tui.NewPadder(1, 0, p.Register),
	))

	p.Root = tui.NewVBox(NewContent(wBox), p.status)

	return p
}
