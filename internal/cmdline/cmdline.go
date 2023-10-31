package cmdline

import (
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/component"
	"github.com/ravsii/minigun/internal/screen"
)

var (
	cursorInput = tcell.StyleDefault.Underline(true)
)

var _ component.Component = (*CommandLine)(nil)

type CommandLine struct {
}

func New() *CommandLine {
	return &CommandLine{}
}

func (c *CommandLine) Draw() {
	c.DrawInput("", -1)
}

func (c *CommandLine) DrawInput(input string, cursorAt int) {
	c.printStyled(":"+input, tcell.StyleDefault, cursorAt)
}

func (c *CommandLine) HandleUserInput() string {
	input := make([]rune, 0, 20)
	var cursorPos int

inputLoop:
	for {
		c.DrawInput(string(input), cursorPos)
		screen.Show()

		event, ok := screen.Screen().PollEvent().(*tcell.EventKey)
		if !ok {
			continue
		}

		switch {
		case event.Modifiers() != tcell.ModNone:
			break
		case event.Key() == tcell.KeyEnter:
			break inputLoop
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

	return string(input)
}

// printStyled prints the given message using the given style.
// If the cursor is being hidden if the value is negative
func (c *CommandLine) printStyled(msg string, style tcell.Style, cursorAt int) {
	_, y := screen.Screen().Size()
	y--
	screen.Screen().HideCursor()

	if cursorAt >= 0 {
		cursorAt++
	}

	var x int
	for _, r := range msg {
		style := tcell.StyleDefault
		if x == cursorAt {
			screen.Screen().ShowCursor(x, y)
		}
		screen.Screen().SetContent(x, y, r, nil, style)

		x++
	}

	if x == cursorAt {
		screen.Screen().SetContent(x, y, ' ', nil, cursorInput)
		x++
	}

	c.printEmptyFrom(x)
}

func (c *CommandLine) printEmptyFrom(x int) {
	for w, y := screen.Screen().Size(); x < w; x++ {
		screen.Screen().SetContent(x, y-1, ' ', nil, tcell.StyleDefault)
	}
}
