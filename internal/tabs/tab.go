package tabs

import (
	"bytes"
	"io"
	"math"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/component"
	"github.com/ravsii/minigun/internal/screen"
)

type Mode int

const (
	ModeView Mode = iota
	ModeInsert
)

var _ component.Component = (*Tab)(nil)

type Tab struct {
	lines  []string
	cursor Cursor

	w       int
	h       int
	xOffset int
	yOffset int
}

func New(w, h, xOffset, yOffset int) *Tab {
	t := &Tab{
		w:       w,
		h:       h,
		xOffset: xOffset,
		yOffset: yOffset,
	}

	return t
}

func (t *Tab) Resize() {
	w, h := screen.Screen().Size()
	h -= 2
	t.w = w
	t.h = h
}

func (t *Tab) FromPath(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	lineBytes := bytes.Split(b, []byte("\n"))
	lines := make([]string, len(lineBytes))
	for i := range lineBytes {
		lines[i] = string(lineBytes[i])
	}

	t.lines = lines
	t.Draw()
	return nil
}

func (t *Tab) Draw() {
	xx, yy := t.w, t.h

	// amount of cells we need to store the biggest line number
	xStart := t.lineNumbersWidth() + 2

	var start, end, cursorLine int

	if len(t.lines) != 0 {
		// todo: this needs to be refactored

		half := int(math.Ceil(float64(yy) / 2))
		end = max(t.cursor.Line+half, yy)
		lmh := t.cursor.Line - half
		start = max(lmh, 0)

		// we can't scroll more than the amount of line in the file
		if end+half > len(t.lines) {
			end = len(t.lines)
		}

		// at the enf of the file, start should be equals to
		// last line number - term height
		if start > len(t.lines)/2 && start > end-yy {
			start = end - yy
		}

		cursorLine = t.cursor.Line - start
	} else {
		start, end, cursorLine = 0, 0, 0
	}

	t.drawLineNumbers(start)

	for y, line := range t.lines[start:end] {
		if y >= t.h {
			break
		}

		for x, c := range line {
			style := tcell.StyleDefault
			if cursorLine == y && t.cursor.Position == x {
				screen.Screen().ShowCursor(x+xStart, y)
				style = cursorStyle
			}

			switch c {
			case '\n':
				screen.Screen().SetContent(x+xStart, y, c, nil, cursorStyle)
			default:
				screen.Screen().SetContent(x+xStart, y, c, nil, style)
			}
		}

		if len(line) == 1 && cursorLine == y {
			screen.Screen().ShowCursor(3, y)
		}

		if cursorLine == y && xx > len(line) {
			for x := len(line); x < xx; x++ {
				screen.Screen().SetContent(x+xStart, y, ' ', nil, cursorStyle)
			}
		} else {
			for x := len(line); x < xx; x++ {
				screen.Screen().SetContent(x+xStart, y, ' ', nil, tcell.StyleDefault)
			}

		}
	}

	if t.h > end {
		for y := end; y < t.h; y++ {
			screen.FillLineEmpty(y, tcell.StyleDefault)
		}
	}
}

func (t *Tab) Cursor() *Cursor {
	return &t.cursor
}

func (t *Tab) MoveUp() {
	if t.cursor.Line == 0 {
		return
	}
	t.cursor.Line--
	t.updateCursorPosition()
	t.Draw()
}

func (t *Tab) MoveDown() {
	if t.cursor.Line >= len(t.lines)-1 {
		return
	}
	t.cursor.Line++
	t.updateCursorPosition()
	t.Draw()
}

// updateCursorPosition updates cursor position on the current line.
//   - If current line is long enough to use cursor's PrevPosition,
func (t *Tab) updateCursorPosition() {
	newLineLen := len(t.lines[t.cursor.Line]) - 1

	if t.cursor.PrevPosition <= newLineLen {
		t.cursor.Position = t.cursor.PrevPosition
	} else {
		t.cursor.Position = max(newLineLen, 0)
	}
}

func (t *Tab) MoveLeft() {
	if t.cursor.Position == 0 {
		return
	}
	t.cursor.Position--
	t.cursor.PrevPosition = t.cursor.Position
	t.Draw()
}

func (t *Tab) MoveRight() {
	// -1 for \n
	if t.cursor.Position >= len(t.lines[t.cursor.Line])-1 {
		return
	}
	t.cursor.Position++
	t.cursor.PrevPosition = t.cursor.Position
	t.Draw()
}
