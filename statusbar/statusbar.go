package statusbar

import (
	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/mode"
)

var statusBg = tcell.StyleDefault.Background(tcell.ColorGold)

type StatusBar struct {
	s tcell.Screen
}

func New(s tcell.Screen) *StatusBar {
	return &StatusBar{
		s: s,
	}
}

func (s *StatusBar) Draw(y int, w int) {
	modeStr := mode.String()
	modeColor := mode.Color()

	style := tcell.StyleDefault.Background(modeColor)

	l := len(modeStr)

	for x := 0; x < 4; x++ {
		if x < 2 {
			s.s.SetContent(x, y, ' ', nil, style)
		} else {
			s.s.SetContent(x+l, y, ' ', nil, style)
		}
	}

	for x, r := range modeStr {
		s.s.SetContent(x+2, y, r, nil, style)
	}
	for x := l + 4; x < w; x++ {
		s.s.SetContent(x, y, ' ', nil, statusBg)
	}
}
