package view

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/registry"
)

type MainPage struct {
	PageHooks
	serviceRegistry registry.ClientServiceRegistry

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

func NewMainPage(serviceRegistry registry.ClientServiceRegistry) *MainPage {
	p := &MainPage{
		serviceRegistry: serviceRegistry,
		SaveData:        tui.NewButton("[Сохранить данные]"),
		LoadData:        tui.NewButton("[Получить данные]"),
		welcomeLabel:    tui.NewLabel(""),
	}

	box := tui.NewVBox(
		p.SaveData,
		p.LoadData,
	)
	box.SetFocused(true)

	wBlock := NewWindowBlockLabel(p.welcomeLabel)
	wBlock.Append(tui.NewSpacer())
	wBlock.Append(box)
	wBlock.Append(tui.NewSpacer())

	p.Root = tui.NewVBox(NewContent(wBlock))

	return p
}
