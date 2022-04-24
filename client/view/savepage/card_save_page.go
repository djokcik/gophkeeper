package savepage

import (
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/view"
)

type CardSavePage struct {
	view.PageHooks
	Root *tui.Box

	keyField        *tui.Entry
	cardNumberField *tui.Entry
	yearField       *tui.Entry
	cvvField        *tui.Entry

	Submit *tui.Button
	Back   *tui.Button
}

func (p CardSavePage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.cardNumberField, p.yearField, p.cvvField, p.Submit, p.Back}
}

func (p CardSavePage) GetRoot() tui.Widget {
	return p.Root
}

func (p CardSavePage) OnActivated(fn func(b *tui.Button)) {
	for _, button := range []*tui.Button{p.Back, p.Submit} {
		if button == p.Submit {
			if p.cardNumberField.Text() == "" || p.yearField.Text() == "" {
				return
			}

			continue
		}

		button.OnActivated(func(b *tui.Button) { fn(b) })
	}
}

func NewCardSavePage() *CardSavePage {
	p := &CardSavePage{Back: view.NewBackButton()}

	keyField, keyBlock := view.NewEditBlock("Ключ")
	p.keyField = keyField
	p.keyField.SetFocused(true)

	cardNumberField, cardNumberBlock := view.NewEditBlock("Номер карты")
	p.cardNumberField = cardNumberField

	yearField, yearBlock := view.NewEditBlock("Год выпуска")
	p.yearField = yearField

	cvvField, cvvBlock := view.NewEditBlock("cvv")
	p.cvvField = cvvField
	p.cvvField.SetEchoMode(tui.EchoModePassword)

	submit := tui.NewButton("[Сохранить]")
	p.Submit = submit
	p.Back = view.NewBackButton()

	buttons := tui.NewHBox(
		tui.NewPadder(1, 0, p.Back),
		tui.NewSpacer(),
		tui.NewPadder(1, 0, p.Submit),
	)

	window := tui.NewVBox(
		tui.NewPadder(10, 0, tui.NewLabel(view.Logo)),
		tui.NewSpacer(),
		tui.NewPadder(1, 0, tui.NewLabel("Укажите номер карты, год выпуска и cvv код")),
		tui.NewLabel(""),
		keyBlock,
		tui.NewSpacer(),
		cardNumberBlock,
		tui.NewSpacer(),
		yearBlock,
		tui.NewSpacer(),
		cvvBlock,
		tui.NewLabel(""),
		buttons,
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
