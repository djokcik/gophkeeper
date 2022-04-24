package savepage

import (
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/view"
)

type TextDataSavePage struct {
	view.PageHooks
	Root *tui.Box

	keyField      *tui.Entry
	textDataField *tui.Entry

	Submit *tui.Button
	Back   *tui.Button
}

func (p TextDataSavePage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.textDataField, p.Submit, p.Back}
}

func (p TextDataSavePage) GetRoot() tui.Widget {
	return p.Root
}

func (p TextDataSavePage) OnActivated(fn func(b *tui.Button)) {
	for _, button := range []*tui.Button{p.Back, p.Submit} {
		if button == p.Submit {
			if p.keyField.Text() == "" || p.textDataField.Text() == "" {
				return
			}

			continue
		}

		button.OnActivated(func(b *tui.Button) { fn(b) })
	}
}

func NewTextDataSavePage() *TextDataSavePage {
	p := &TextDataSavePage{Back: view.NewBackButton()}

	keyField, keyBlock := view.NewEditBlock("Ключ")
	p.keyField = keyField
	p.keyField.SetFocused(true)

	textDataField, textDataBlock := view.NewEditBlock("Текст")
	p.textDataField = textDataField

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
		tui.NewLabel("Выберите ключ и текстовую информацию"),
		keyBlock,
		tui.NewSpacer(),
		textDataBlock,
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
