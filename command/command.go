package command

import (
	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/mode"
	"github.com/ravsii/minigun/tab"
)

var (
	cursorInput = tcell.StyleDefault.Blink(true).Underline(true)
)

type Command struct {
	s tcell.Screen
}

func New(s tcell.Screen) *Command {
	return &Command{
		s: s,
	}
}

func (c *Command) Draw() {
	c.DrawInput("", false)
}

func (c *Command) DrawInput(input string, typing bool) {
	w, y := c.s.Size()
	y--

	c.s.SetContent(0, y, ':', nil, tcell.StyleDefault)
	x := 1
	for _, r := range input {
		c.s.SetContent(x+1, y, r, nil, tcell.StyleDefault)
		x++
	}
	for x := len(input); x < w; x++ {
		c.s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
	}
	if typing {
		c.s.SetContent(len(input)+1, y, ' ', nil, cursorInput)
	}
}

func (c *Command) HandleInput() {
	mode.Set(mode.Console)
	defer mode.Set(mode.View)

	t := tab.Root().ActiveTab()

	var b string
outer:
	for {
		event, ok := c.s.PollEvent().(*tcell.EventKey)
		if !ok {
			continue
		}

		switch {
		case event.Modifiers() != tcell.ModNone:
			break
		case event.Key() == tcell.KeyEnter:
			break outer
		case event.Key() == tcell.KeyBackspace || event.Key() == tcell.KeyBackspace2:
			if len(b) > 0 {
				b = b[:len(b)-1]
			}
		default:
			b += string(event.Rune())
			c.DrawInput(b, true)
			t.Draw()
			c.s.Show()
		}
	}

}
