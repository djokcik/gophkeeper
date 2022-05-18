package view

import (
	"github.com/marcusolsson/tui-go"
	"log"
)

type UIClient struct {
	UI     tui.UI
	widget Page
}

func NewUIClient() *UIClient {
	ui, err := tui.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	return &UIClient{
		UI: ui,
	}
}

func (c *UIClient) Start() {
	c.ResetKeybinding()

	if err := c.UI.Run(); err != nil {
		log.Fatal(err)
	}
}

func (c *UIClient) ResetKeybinding() {
	c.UI.ClearKeybindings()
	c.UI.SetKeybinding("Esc", func() { c.UI.Quit() })
}

func (c *UIClient) SetWidget(widget Page) {
	focusChain := &tui.SimpleFocusChain{}
	focusChain.Set(widget.GetFocusChain()...)

	c.UI.SetFocusChain(focusChain)

	if c.widget != nil {
		c.widget.After()
	}

	c.ResetKeybinding()
	widget.Before()
	c.UI.SetWidget(widget.GetRoot())
	c.widget = widget
}
