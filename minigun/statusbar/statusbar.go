package statusbar

import (
	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/minigun/mode"
)

var (
	statusBg  = tcell.StyleDefault.Background(tcell.ColorGold)
	statusBar *StatusBar
)

type StatusBar struct {
	s tcell.Screen
}

func Init(s tcell.Screen) *StatusBar {
	if statusBar == nil {
		statusBar = &StatusBar{s: s}
	}

	return statusBar
}

// Get returns an instance of a Command struct.
func Get() *StatusBar {
	if statusBar == nil {
		panic("Init() is not called")
	}

	return statusBar
}
func (s *StatusBar) Draw() {
	w, h := s.s.Size()
	y := h - 2

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
