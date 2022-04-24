package savepage

import (
	"github.com/marcusolsson/tui-go"
	"gophkeeper/client/view"
)

type BinaryDataSavePage struct {
	view.PageHooks
	Root *tui.Box

	keyField        *tui.Entry
	pathToFileField *tui.Entry

	Submit *tui.Button
	Back   *tui.Button
}

func (p BinaryDataSavePage) GetFocusChain() []tui.Widget {
	return []tui.Widget{p.keyField, p.pathToFileField, p.Submit, p.Back}
}

func (p BinaryDataSavePage) GetRoot() tui.Widget {
	return p.Root
}

func (p BinaryDataSavePage) OnActivated(fn func(b *tui.Button)) {
	for _, button := range []*tui.Button{p.Back, p.Submit} {
		if button == p.Submit {
			if p.keyField.Text() == "" || p.pathToFileField.Text() == "" {
				return
			}

			continue
		}

		button.OnActivated(func(b *tui.Button) { fn(b) })
	}
}

func NewBinaryDataSavePage() *BinaryDataSavePage {
	p := &BinaryDataSavePage{Back: view.NewBackButton()}

	p.keyField = tui.NewEntry()
	p.keyField.SetFocused(true)
	formKey := tui.NewGrid(0, 0)
	formKey.AppendRow(p.keyField)
	keyBlock := tui.NewHBox(formKey)
	keyBlock.SetTitle("Ключ")
	keyBlock.SetBorder(true)

	p.pathToFileField = tui.NewEntry()
	formKey = tui.NewGrid(0, 0)
	formKey.AppendRow(p.pathToFileField)
	textDataBlock := tui.NewHBox(formKey)
	textDataBlock.SetTitle("Путь до файла")
	textDataBlock.SetBorder(true)

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
		tui.NewLabel("Выберите ключ и путь до файла"),
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
