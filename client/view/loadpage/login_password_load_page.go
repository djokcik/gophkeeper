package loadpage

import (
	"context"
	"fmt"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/registry"
	"gophkeeper/client/view"
	"gophkeeper/pkg/logging"
)

type LoginPasswordLoadPage struct {
	view.PageHooks
	serviceRegistry registry.ClientServiceRegistry

	Root *tui.Box

	loginField *tui.Entry
	result     *tui.Label
	status     *tui.StatusBar

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
	service := p.serviceRegistry.GetLoginPasswordService()

	for _, button := range []*tui.Button{p.Submit, p.Back} {
		button.OnActivated(func(b *tui.Button) {
			p.status.SetText("Загрузка...")

			ctx := context.Background()
			log := logging.NewFileLogger()

			if b == p.Submit {
				if p.loginField.Text() == "" {
					return
				}

				data, err := service.LoadPasswordByLogin(ctx, p.loginField.Text())
				if err != nil {
					p.status.SetText(fmt.Sprintf("Error: %s", err.Error()))
					log.Error().Err(err).Msg("Submit: invalid LoadPasswordByLogin")

					return
				}

				p.status.SetText("Пароль успешно получен")
				p.result.SetText(fmt.Sprintf("Пароль: %s", data.Password))

				return
			}

			fn(b)
		})
	}
}

func NewLoginPasswordLoadPage(serviceRegistry registry.ClientServiceRegistry) *LoginPasswordLoadPage {
	p := &LoginPasswordLoadPage{
		Back:            view.NewBackButton(),
		serviceRegistry: serviceRegistry,
		status:          tui.NewStatusBar(""),
	}

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

	p.Root = tui.NewVBox(content, p.status)

	return p
}
