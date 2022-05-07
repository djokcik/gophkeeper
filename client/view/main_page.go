package view

import (
	"context"
	"fmt"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/registry"
	"gophkeeper/pkg/logging"
)

type MainPage struct {
	PageHooks
	serviceRegistry registry.ClientServiceRegistry

	Root *tui.Box

	SaveData *tui.Button
	LoadData *tui.Button
	SyncData *tui.Button
	Table    *tui.Table

	welcomeLabel *tui.Label
	status       *tui.StatusBar
}

func (p MainPage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.SaveData, p.LoadData, p.SyncData}
}

func (p MainPage) GetRoot() tui.Widget {
	return p.Root
}

func (p MainPage) OnActivated(fn func(*tui.Button)) {
	for _, button := range []*tui.Button{p.SaveData, p.LoadData} {
		button.OnActivated(func(b *tui.Button) { fn(b) })
	}

	p.SyncData.OnActivated(func(b *tui.Button) {
		ctx := context.Background()
		log := logging.NewFileLogger()

		service := p.serviceRegistry.GetStorageService()

		actions, err := service.LoadRecords(ctx)
		if err != nil {
			p.status.SetText(err.Error())
			log.Error().Err(err).Msg("OnActivated: invalid load records")
			return
		}

		if len(actions) == 0 {
			p.status.SetText("Данные с сервером синхронизированы")
			return
		}

		err = service.SyncServer(ctx)
		if err != nil {
			p.status.SetText(err.Error())
			return
		}

		p.status.SetText("Данные с сервером успешно синхронизированы")
	})
}

func (p *MainPage) Before() {
	user := p.serviceRegistry.GetUserService().GetUser()
	if user.Token == "" {
		p.welcomeLabel.SetText(fmt.Sprintf("\nДобрый день, %s. Сервер недоступен. Переход в offline режим\n", user.Username))
	} else {
		p.welcomeLabel.SetText(fmt.Sprintf("\nДобрый день, %s. (Используйте TAB для навигации)\n", user.Username))
	}
}

func NewMainPage(serviceRegistry registry.ClientServiceRegistry) *MainPage {
	p := &MainPage{
		serviceRegistry: serviceRegistry,
		SaveData:        tui.NewButton("[Сохранить данные]"),
		LoadData:        tui.NewButton("[Получить данные]"),
		SyncData:        tui.NewButton("[Синхронизировать]"),
		welcomeLabel:    tui.NewLabel(""),
		status:          NewStatusLabel(),
	}

	box := tui.NewVBox(
		p.SaveData,
		p.LoadData,
		tui.NewLabel(""),
		p.SyncData,
	)
	box.SetFocused(true)

	wBlock := NewWindowBlockLabel(p.welcomeLabel)
	wBlock.Append(tui.NewSpacer())
	wBlock.Append(box)
	wBlock.Append(tui.NewSpacer())

	p.Root = tui.NewVBox(NewContent(wBlock), p.status)

	return p
}
