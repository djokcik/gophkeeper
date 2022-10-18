package view

import "github.com/marcusolsson/tui-go"

// PageHooks base webhooks
type PageHooks struct {
}

// Before hooks run before widget started
func (p *PageHooks) Before() {
}

// After hooks run after widget destroy
func (p *PageHooks) After() {
}

// Page wedget interface with methods
type Page interface {
	Before()
	After()
	GetFocusChain() []tui.Widget
	GetRoot() tui.Widget
	OnActivated(fn func(b *tui.Button))
}
