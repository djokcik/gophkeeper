package view

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/registry"
)

type MainPage struct {
	PageHooks
	serviceRegistry registry.ServiceRegistry

	Root *tui.Box

	SaveData *tui.Button
	LoadData *tui.Button
	Table    *tui.Table

	welcomeLabel *tui.Label
}

func (p MainPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.SaveData, p.LoadData}
}

func (p MainPage) GetRoot() tui.Widget {
	return p.Root
}

func (p MainPage) OnActivated(fn func(*tui.Button)) {
	for _, button := range []*tui.Button{p.SaveData, p.LoadData} {
		button.OnActivated(func(b *tui.Button) { fn(b) })
	}
}

func (p *MainPage) Before() {
	name := p.serviceRegistry.GetUserService().GetUser().Username
	p.welcomeLabel.SetText(fmt.Sprintf("\nДобрый день, %s. (Используйте TAB для навигации)", name))
}

func NewMainPage(serviceRegistry registry.ServiceRegistry) *MainPage {
	p := &MainPage{serviceRegistry: serviceRegistry}

	p.SaveData = tui.NewButton("[Сохранить данные]")
	p.LoadData = tui.NewButton("[Получить данные]")
	p.welcomeLabel = tui.NewLabel("")

	box := tui.NewVBox(
		p.SaveData,
		p.LoadData,
	)
	box.SetFocused(true)

	window := tui.NewVBox(
		tui.NewPadder(10, 0, tui.NewLabel(Logo)),
		tui.NewSpacer(),
		p.welcomeLabel,
		tui.NewLabel(""),
		box,
		tui.NewLabel(""),
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
