package controller

import (
	"context"
	"github.com/djokcik/gophkeeper/client/registry"
	"github.com/djokcik/gophkeeper/client/view"
	"github.com/marcusolsson/tui-go"
)

// Start starts terminal user interface
func Start(_ context.Context, serviceRegistry registry.ClientServiceRegistry) {
	client := view.NewUIClient()

	ctr := NewUIController(serviceRegistry)
	router := GenerateRouter(ctr)

	ctr.RegisterLoginPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.MainPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.SavePage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.LoadPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.RemovePage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.LoginPasswordSavePage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.BinaryDataSavePage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.TextDataSavePage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.CardSavePage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.LoginPasswordLoadPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.TextDataLoadPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.BinaryDataLoadPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.BankCardLoadPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })

	client.SetWidget(ctr.RegisterLoginPage)

	client.Start()
}
