package statusbar

import (
	"fmt"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/components"
	"github.com/ravsii/minigun/internal/components/box"
	"github.com/ravsii/minigun/internal/components/textbox"
	"github.com/ravsii/minigun/internal/mode"
	"github.com/ravsii/minigun/internal/screen"
)

var modeColor = map[mode.Mode]tcell.Color{
	mode.View:    tcell.NewHexColor(0xFF0000),
	mode.Command: tcell.NewHexColor(0x00FF00),
	mode.Replace: tcell.NewHexColor(0x000FF0),
}

var _ components.Component = (*StatusBar)(nil)

type StatusBar struct {
	cursorLine, cursorPos int
}

func New() *StatusBar {
	return &StatusBar{}
}

func (s *StatusBar) SetCursor(l, p int) {
	s.cursorLine, s.cursorPos = l+1, p+1
}

func (s *StatusBar) Draw() {
	w, h := screen.Screen().Size()
	y := h - 2

	screen.FillLineEmpty(y, tcell.StyleDefault.Background(tcell.ColorGold))

	modeStr := mode.Current().String()
	modeColor := modeColor[mode.Current()]

	paddingY2 := components.Padding{Left: 2, Right: 2}

	modeTextBox := textbox.New(0, y, modeStr,
		box.WithTextColor(tcell.ColorBlack),
		box.WithBackground(modeColor),
		box.WithPadding(paddingY2))

	cursorString := fmt.Sprintf("Line: %d, Col: %d", s.cursorLine, s.cursorPos)
	lenStr := utf8.RuneCountInString(cursorString)
	cursorTbX := w - paddingY2.SumX() - lenStr
	cursorTextBox := textbox.New(cursorTbX, y, cursorString,
		box.WithTextColor(tcell.ColorBlack),
		box.WithBackground(tcell.ColorPurple),
		box.WithPadding(paddingY2))

	modeTextBox.Draw()
	cursorTextBox.Draw()
}
