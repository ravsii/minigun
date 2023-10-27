package cmd

import (
	"strings"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/minigun/mode"
	"github.com/ravsii/minigun/minigun/statusbar"
)

const (
	block     = '█'
	bar       = '|'
	underline = '_'
)

var (
	command     *Command
	cursorInput = tcell.StyleDefault.Underline(true)
)

type Command struct {
	s tcell.Screen

	errMsg string
}

func Init(s tcell.Screen) *Command {
	if command == nil {
		command = &Command{s: s}
	}

	return command
}

// Get returns an instance of a Command struct.
func Get() *Command {
	if command == nil {
		panic("Init() is not called")
	}

	return command
}

func (c *Command) Draw() {
	if c.errMsg != "" {
		return
	}
	c.DrawInput("", -1)
}

func (c *Command) DrawInput(input string, cursorAt int) {
	c.printStyled(":"+input, tcell.StyleDefault, cursorAt)
}

func (c *Command) HandleInput() {
	mode.Set(mode.Console)
	statusbar.Get().Draw()
	c.s.SetCursorStyle(tcell.CursorStyleBlinkingBar)
	defer func() {
		mode.Set(mode.View)
		statusbar.Get().Draw()
		c.s.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
		c.s.HideCursor()
	}()

	userCommand := c.handleInputString()
	splitted := strings.Split(userCommand, " ")
	Execute(splitted...)
}

func (c *Command) handleInputString() string {
	input := make([]rune, 0, 20)
	var cursorPos int

	// todo: this needs to be refactored

	for {
		c.DrawInput(string(input), cursorPos)
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
		case event.Key() == tcell.KeyHome:
			cursorPos = 0
		case event.Key() == tcell.KeyEnd:
			cursorPos = len(input)
		default:
			r := event.Rune()
			if !unicode.IsGraphic(r) {
				break
			}

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

// printStyled prints the given message using the given style.
// If the cursor is being hidden if the value is negative
func (c *Command) printStyled(msg string, style tcell.Style, cursorAt int) {
	_, y := c.s.Size()
	y--

	if cursorAt >= 0 {
		cursorAt++
	}

	var x int
	for _, r := range msg {
		style := tcell.StyleDefault
		if x == cursorAt {
			c.s.ShowCursor(x, y)
		}
		c.s.SetContent(x, y, r, nil, style)

		x++
	}

	if x == cursorAt {
		c.s.SetContent(x, y, ' ', nil, cursorInput)
		x++
	}

	c.printEmptyFrom(x)
	c.s.Show()
}

func (c *Command) printEmptyFrom(x int) {
	for w, y := c.s.Size(); x < w; x++ {
		c.s.SetContent(x, y-1, ' ', nil, tcell.StyleDefault)
	}
}
