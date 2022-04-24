package controller

import (
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/view"
)

func Start() {
	client := view.NewUiClient()

	ctr := NewUIController()
	router := GenerateRouter(ctr)

	ctr.RegisterLoginPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.MainPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.SavePage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })
	ctr.LoadPage.OnActivated(func(b *tui.Button) { client.SetWidget(router[b]) })

	client.SetWidget(ctr.RegisterLoginPage)

	client.Start()
}
