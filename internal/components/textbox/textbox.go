package textbox

import (
	"unicode/utf8"

	"github.com/ravsii/minigun/internal/components/box"
)

type TextBox struct {
	text []rune
	box.Box
}

func New(x, y int, text string, opts ...box.BoxOption) *TextBox {
	w := utf8.RuneCountInString(text)
	textbox := TextBox{
		Box:  (*box.New(x, y, w, 1, opts...)),
		text: []rune(text),
	}

	return &textbox
}

func (tb *TextBox) Draw() {
	tb.Box.Draw(func(x, _ int) rune {
		if x > len(tb.text)-1 {
			return ' '
		}

		return tb.text[x]
	})
}
