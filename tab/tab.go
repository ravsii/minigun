package tab

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/statusbar"
)

type Mode int

const (
	ModeView Mode = iota
	ModeInsert
)

type Tab struct {
	s tcell.Screen

	Lines  []string
	Cursor Cursor

	w       int
	h       int
	xOffset int
	yOffset int

	sb *statusbar.StatusBar

	parent *Group
}

func newTab(s tcell.Screen, w, h, xOffset, yOffset int, parent *Group) *Tab {
	sb := statusbar.New(s)

	return &Tab{
		s:       s,
		w:       w,
		h:       h,
		xOffset: xOffset,
		yOffset: yOffset,
		sb:      sb,
		parent:  parent,
	}
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

	t.Lines = lines
	return nil
}

func (t *Tab) Draw() {
	xx, yy := t.w, t.h

	// todo: this needs to be refactored

	half := int(math.Ceil(float64(yy) / 2))
	end := max(t.Cursor.Line+half, yy)
	lmh := t.Cursor.Line - half
	start := max(lmh, 0)

	// we can't scroll more than the amount of line in the file
	if end+half > len(t.Lines) {
		end = len(t.Lines)
	}

	// at the enf of the file, start should be equals to
	// last line number - term height
	if start > len(t.Lines)/2 && start > end-yy {
		start = end - yy
	}

	cursorLine := t.Cursor.Line - start

	for y, line := range t.Lines[start:end] {
		lineStr := make([]rune, 3)
		st := fmt.Sprint(start + y + 1)
		for i, r := range st {
			lineStr[i] = r
		}
		for i, r := range lineStr {
			t.s.SetContent(i, y, r, nil, lineNumberStyle)
		}

		// TODO: fix no cursor on empty line

		for x, c := range line {
			style := tcell.StyleDefault
			if cursorLine == y && t.Cursor.Position == x {
				t.s.ShowCursor(x+3, y)
				style = cursorStyle
			}

			switch c {
			case '\n':
				t.s.SetContent(x+3, y, c, nil, cursorStyle)
			default:
				t.s.SetContent(x+3, y, c, nil, style)
			}
		}

		if len(line) == 1 && cursorLine == y {
			t.s.ShowCursor(3, y)
		}

		if cursorLine == y && xx > len(line) {
			for x := len(line); x < xx; x++ {
				t.s.SetContent(x+3, y, ' ', nil, cursorStyle)
			}
		} else {
			for x := len(line); x < xx; x++ {
				t.s.SetContent(x+3, y, ' ', nil, tcell.StyleDefault)
			}

		}
	}

	t.sb.Draw(t.yOffset+t.h, t.w)
}

// HandleKey handles key event. It returns true if a user desires to quit the app.
func (t *Tab) HandleKey(s tcell.Screen, key *tcell.EventKey) bool {
	switch {
	case key.Key() == tcell.KeyCtrlC:
		return true
	case key.Key() == tcell.KeyCtrlL:
		s.Sync()
	case key.Rune() == 'C' || key.Rune() == 'c':
		s.Clear()
	case key.Rune() == 'H' || key.Rune() == 'h':
		t.MoveLeft()
	case key.Rune() == 'J' || key.Rune() == 'j':
		t.MoveDown()
	case key.Rune() == 'K' || key.Rune() == 'k':
		t.MoveUp()
	case key.Rune() == 'L' || key.Rune() == 'l':
		t.MoveRight()
	}

	return false
}

func (t *Tab) StatusBar() *statusbar.StatusBar {
	return t.sb
}

func (t *Tab) MoveUp() {
	if t.Cursor.Line == 0 {
		return
	}

	newLine := t.Lines[t.Cursor.Line-1]
	lnl := len(newLine) - 1

	if t.Cursor.prevPosition <= lnl {
		t.Cursor.Position = t.Cursor.prevPosition
	} else {
		t.Cursor.Position = max(lnl, 0)
	}

	t.Cursor.Line--
}

func (t *Tab) MoveDown() {
	if t.Cursor.Line >= len(t.Lines)-1 {
		return
	}

	newLine := t.Lines[t.Cursor.Line+1]
	lnl := len(newLine) - 1

	if t.Cursor.prevPosition <= lnl {
		t.Cursor.Position = t.Cursor.prevPosition
	} else {
		t.Cursor.Position = max(lnl, 0)
	}

	t.Cursor.Line++
}

func (t *Tab) MoveLeft() {
	if t.Cursor.Position == 0 {
		return
	}
	t.Cursor.Position--
	t.Cursor.prevPosition = t.Cursor.Position
}

func (t *Tab) MoveRight() {
	// -1 for \n
	if t.Cursor.Position >= len(t.Lines[t.Cursor.Line])-1 {
		return
	}
	t.Cursor.Position++
	t.Cursor.prevPosition = t.Cursor.Position
}
