package view

import (
	"github.com/marcusolsson/tui-go"
	"log"
)

type Page interface {
	GetFocusChain() []tui.Widget
	GetRoot() *tui.Box
	OnActivated(fn func(b *tui.Button))
	GetButtons() []*tui.Button
}

type UIClient struct {
	UI tui.UI
}

func NewUiClient() *UIClient {
	ui, err := tui.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	return &UIClient{
		UI: ui,
	}
}

func (c *UIClient) Start() {
	c.UI.SetKeybinding("Esc", func() { c.UI.Quit() })

	if err := c.UI.Run(); err != nil {
		log.Fatal(err)
	}
}

func (c *UIClient) SetWidget(widget Page) {
	focusChain := &tui.SimpleFocusChain{}
	focusChain.Set(widget.GetFocusChain()...)

	c.UI.SetFocusChain(focusChain)
	c.UI.SetWidget(widget.GetRoot())
}
