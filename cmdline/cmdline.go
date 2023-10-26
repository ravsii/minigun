package cmdline

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/commands"
	"github.com/ravsii/minigun/mode"
	"github.com/ravsii/minigun/tab"
)

const (
	block     = '█'
	bar       = '|'
	underline = '_'
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
	c.DrawInput("", -1)
}

func (c *Command) DrawInput(input string, showCursorAt int) {
	w, y := c.s.Size()
	y--

	il := len(input)

	c.s.SetContent(0, y, ':', nil, tcell.StyleDefault)
	x := 1
	if showCursorAt == 0 {
		r := bar
		if x == il+1 {
			r = block
		}
		c.s.SetContent(x, y, r, nil, cursorInput)
		x++
	}

	for _, r := range input {
		c.s.SetContent(x, y, r, nil, tcell.StyleDefault)
		if x == showCursorAt {
			x++
			r := bar
			if x == il+1 {
				r = block
			}
			c.s.SetContent(x, y, r, nil, cursorInput)
		}

		x++
	}

	for ; x < w; x++ {
		c.s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
	}
}

func (c *Command) HandleInput() {
	mode.Set(mode.Console)
	defer mode.Set(mode.View)

	userCommand := c.handleInputString()
	splitted := strings.Split(userCommand, " ")
	commands.Execute(splitted...)
}

func (c *Command) handleInputString() string {
	t := tab.Root().ActiveTab()

	input := make([]rune, 0, 20)
	var cursorPos int

	// todo: this needs to be refactored

	for {
		c.DrawInput(string(input), cursorPos)
		t.Draw()
		c.s.Show()

		event, ok := c.s.PollEvent().(*tcell.EventKey)
		if !ok {
			continue
		}

		switch {
		case event.Modifiers() != tcell.ModNone:
			break
		case event.Key() == tcell.KeyEnter:
			return string(input)
		case event.Key() == tcell.KeyBackspace || event.Key() == tcell.KeyBackspace2:
			if len(input) == 0 {
				// vim does quit command mode on empty command & backspace
				return ""
			}

			input = append(input[:cursorPos-1], input[cursorPos:]...)
			cursorPos--
		case event.Key() == tcell.KeyLeft:
			if cursorPos > 0 {
				cursorPos--
			}
		case event.Key() == tcell.KeyRight:
			if cursorPos < len(input) {
				cursorPos++
			}
		default:
			r := event.Rune()
			if cursorPos == len(input) {
				input = append(input, r)
			} else {
				input = append(input[:cursorPos+1], input[cursorPos:]...)
				input[cursorPos] = r
			}

			cursorPos++
		}
	}
}
