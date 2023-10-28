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

func Finish() {
	// You have to catch panics in a defer, clean up, and re-raise them.
	// Otherwise your application can die without leaving any diagnostic trace.
	r := recover()
	s.Fini()
	if r != nil {
		panic(r)
	}
}
