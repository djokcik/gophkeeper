package savepage

import (
	"context"
	"fmt"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/registry"
	"gophkeeper/client/view"
	"gophkeeper/pkg/logging"
)

type LoginPasswordSavePage struct {
	view.PageHooks
	serviceRegistry registry.ClientServiceRegistry

	Root *tui.Box

	loginField    *tui.Entry
	passwordField *tui.Entry
	status        *tui.StatusBar

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
	loginPasswordService := p.serviceRegistry.GetLoginPasswordService()

	p.Back.OnActivated(fn)
	p.Submit.OnActivated(func(b *tui.Button) {
		ctx := context.Background()
		log := logging.NewFileLogger()

		if p.loginField.Text() == "" || p.passwordField.Text() == "" {
			return
		}

		err := loginPasswordService.SaveLoginPassword(ctx, p.loginField.Text(), p.passwordField.Text())
		if err != nil {
			p.status.SetText(fmt.Sprintf("Error: %s", err.Error()))
			log.Error().Err(err).Msg("Submit: invalid SaveLoginPassword")

			return
		}

		p.status.SetText("Данные успешно сохранены")
	})
}

func NewLoginPasswordSagePage(serviceRegistry registry.ClientServiceRegistry) *LoginPasswordSavePage {
	p := &LoginPasswordSavePage{
		Back:            view.NewBackButton(),
		status:          tui.NewStatusBar(""),
		serviceRegistry: serviceRegistry,
	}

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

	p.Root = tui.NewVBox(content, p.status)

	return p
}
