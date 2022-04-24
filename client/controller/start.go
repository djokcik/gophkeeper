package controller

import (
	"context"
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/registry"
	"gophkeeper/client/view"
)

func Start(_ context.Context, serviceRegistry registry.ServiceRegistry) {
	client := view.NewUiClient()

	ctr := NewUIController(serviceRegistry)
	router := GenerateRouter(ctr)

	ctr.RegisterLoginPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.MainPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.SavePage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.LoadPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.LoginPasswordSavePage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.BinaryDataSavePage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.TextDataSavePage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.CardSavePage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })

	client.SetWidget(ctr.RegisterLoginPage)

	client.Start()
}
