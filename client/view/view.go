package view

import "github.com/marcusolsson/tui-go"

type PageHooks struct {
}

func (p *PageHooks) Before() {
}

func (p *PageHooks) After() {
}

type Page interface {
	Before()
	After()
	GetFocusChain() []tui.Widget
	GetRoot() tui.Widget
	OnActivated(fn func(b *tui.Button))
}
