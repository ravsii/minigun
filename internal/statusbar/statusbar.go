package statusbar

import (
	"fmt"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/component"
	"github.com/ravsii/minigun/internal/component/box"
	"github.com/ravsii/minigun/internal/component/textbox"
	"github.com/ravsii/minigun/internal/mode"
	"github.com/ravsii/minigun/internal/screen"
)

var modeColor = map[mode.Mode]tcell.Color{
	mode.View:    tcell.NewHexColor(0xFF0000),
	mode.Console: tcell.NewHexColor(0x00FF00),
}

var _ component.Component = (*StatusBar)(nil)

type StatusBar struct {
	cursorLine, cursorPos int
}

func New() StatusBar {
	return StatusBar{}
}

func (s *StatusBar) SetCursor(l, p int) {
	s.cursorLine, s.cursorPos = l+1, p+1
}

func (s *StatusBar) Draw() {
	w, h := screen.Screen().Size()
	y := h - 2

	screen.FillLine(y, tcell.StyleDefault.Background(tcell.ColorGold))

	modeStr := mode.String()
	modeColor := modeColor[mode.Current()]

	paddingY2 := component.Padding{Left: 2, Right: 2}

	modeTextBox := textbox.New(0, y, modeStr, box.WithBackground(modeColor), box.WithPadding(paddingY2))

	cursorString := fmt.Sprintf("Line: %d, Col: %d", s.cursorLine, s.cursorPos)
	lenStr := utf8.RuneCountInString(cursorString)
	cursorTbX := w - paddingY2.SumX() - lenStr
	cursorTextBox := textbox.New(cursorTbX, y, cursorString, box.WithBackground(tcell.ColorPurple), box.WithPadding(paddingY2))

	modeTextBox.Draw()
	cursorTextBox.Draw()
}
