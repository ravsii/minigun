// screen is a wrapper around tcell.Screen for easier use
package screen

import "github.com/gdamore/tcell/v2"

const (
	block     = '█'
	bar       = '|'
	underline = '_'
)

var s tcell.Screen

func init() {
	var err error
	s, err = tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := s.Init(); err != nil {
		panic(err)
	}

	s.Clear()
}

func Screen() tcell.Screen {
	return s
}

func Show() {
	s.Show()
}

func SetEmpty(x, y int) {
	SetEmptyStyle(x, y, tcell.StyleDefault)
}

func SetEmptyStyle(x, y int, style tcell.Style) {
	SetRuneStyle(x, y, ' ', style)
}

func SetRune(x, y int, r rune) {
	SetRuneStyle(x, y, r, tcell.StyleDefault)
}

func SetRuneStyle(x, y int, r rune, style tcell.Style) {
	s.SetContent(x, y, r, nil, style)
}

func FillLine(y int, style tcell.Style) {
	w, _ := s.Size()
	for x := 0; x < w; x++ {
		SetRuneStyle(x, y, ' ', style)
	}
}

func Finish() {
	// You have to catch panics in a defer, clean up, and re-raise them.
	// Otherwise your application can die without leaving any diagnostic trace.
	r := recover()
	s.Clear()
	s.Fini()
	if r != nil {
		panic(r)
	}
}
