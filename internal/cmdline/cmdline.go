package cmdline

import (
	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/components"
	"github.com/ravsii/minigun/internal/screen"
)

const inputSize = 20

var (
	cursorInput = tcell.StyleDefault.Underline(true)
)

var _ components.Component = (*CommandLine)(nil)

type CommandLine struct {
	input     []rune
	cursorPos int
}

func New() *CommandLine {
	return &CommandLine{
		cursorPos: -1,
		input:     make([]rune, 0, inputSize),
	}
}

func (c *CommandLine) Draw() {
	c.printStyled(":"+string(c.input), tcell.StyleDefault, c.cursorPos)
}

// AddRune adds a rune at a current input position
func (c *CommandLine) AddRune(r rune) {
	if (c.cursorPos) < 0 {
		c.cursorPos = 0
	}

	if c.cursorPos == len(c.input) {
		c.input = append(c.input, r)
	} else {
		c.input = append(c.input[:c.cursorPos+1], c.input[c.cursorPos:]...)
		c.input[c.cursorPos] = r
	}

	c.cursorPos++
	c.Draw()
}

// Input retruns current user input. If you don't need the command anymore,
// use c.Reset.
func (c *CommandLine) Input() string {
	return string(c.input)
}

func (c *CommandLine) MoveLeft() {
	if c.cursorPos > 0 {
		c.cursorPos--
	}
	c.Draw()
}

func (c *CommandLine) MoveRight() {
	if c.cursorPos < len(c.input) {
		c.cursorPos++
	}
	c.Draw()
}

func (c *CommandLine) JumpLineStart() {
	c.cursorPos = 0
	c.Draw()
}

func (c *CommandLine) JumpLineEnd() {
	c.cursorPos = len(c.input)
	c.Draw()
}

// DeleteRune removed a current rune before the cursor.
func (c *CommandLine) DeleteRune() {
	if len(c.input) == 0 {
		return
	}

	c.input = append(c.input[:c.cursorPos-1], c.input[c.cursorPos:]...)
	c.cursorPos--
	c.Draw()
}

func (c *CommandLine) Reset() {
	c.input = make([]rune, 0, inputSize)
	c.cursorPos = -1
	c.Draw()
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
