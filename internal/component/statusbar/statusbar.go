package statusbar

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/component"
	"github.com/ravsii/minigun/internal/mode"
	"github.com/ravsii/minigun/internal/screen"
)

var (
	statusBg = tcell.StyleDefault.Background(tcell.ColorGold)
	cursorBg = tcell.StyleDefault.Background(tcell.ColorBlue)
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
	s.cursorLine, s.cursorPos = l, p
}

func (s *StatusBar) Draw() {
	w, h := screen.Screen().Size()
	y := h - 2

	modeStr := mode.String()
	modeColor := modeColor[mode.Current()]

	style := tcell.StyleDefault.Background(modeColor)

	l := len(modeStr)

	for x := 0; x < 4; x++ {
		if x < 2 {
			screen.Screen().SetContent(x, y, ' ', nil, style)
		} else {
			screen.Screen().SetContent(x+l, y, ' ', nil, style)
		}
	}

	for x, r := range modeStr {
		screen.Screen().SetContent(x+2, y, r, nil, style)
	}

	for x := l + 4; x < w; x++ {
		screen.Screen().SetContent(x, y, ' ', nil, statusBg)
	}

	cursorStr := []rune(fmt.Sprintf("Line: %d, Col: %d", s.cursorLine, s.cursorPos))

	i := 0
	for x := w - len(cursorStr); x < w; x++ {
		screen.Screen().SetContent(x, y, cursorStr[i], nil, cursorBg)
		i++
	}

	screen.Screen().Show()
}
