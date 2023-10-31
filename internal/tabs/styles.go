package tabs

import "github.com/gdamore/tcell/v2"

var (
	cursorStyle     = tcell.StyleDefault.Background(tcell.ColorDarkGray)
	cursorStyleCell = cursorStyle.Blink(true).Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)
	lineNumberStyle = tcell.StyleDefault.Background(tcell.ColorAliceBlue).Foreground(tcell.ColorBlack.TrueColor())
)
