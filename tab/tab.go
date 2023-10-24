package tab

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/gdamore/tcell/v2"
)

type Mode int

const (
	ModeView Mode = iota
	ModeInsert
)

type Tab struct {
	Lines  []string
	Cursor Cursor
}

type Cursor struct {
	Line     int
	Position int
	// prevPosition is used to keep horizontal positing when moving up and down
	// between the lines of different width.
	prevPosition int
}

func FromPath(path string) (*Tab, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	lineBytes := bytes.Split(b, []byte("\n"))
	lines := make([]string, len(lineBytes))
	for i := range lineBytes {
		lines[i] = string(lineBytes[i])
	}

	return &Tab{
		Lines: lines,
		Cursor: Cursor{
			Line:     0,
			Position: 0,
		},
	}, nil
}

func (t *Tab) Draw(s tcell.Screen) {
	xx, yy := s.Size()

	// todo: this needs to be refactored

	half := yy / 2
	end := max(t.Cursor.Line+half, yy)
	lmh := t.Cursor.Line - half
	start := max(lmh, 0)
	if end+half > len(t.Lines) {
		end = len(t.Lines)
	}
	if start > len(t.Lines)/2 && start > end-yy {
		start = end - yy
	}
	activeLine := t.Cursor.Line - start
	// fmt.Println(start)

	for y, line := range t.Lines[start:end] {
		lineStr := make([]rune, 3)
		st := fmt.Sprint(start + y + 1)
		for i, r := range st {
			lineStr[i] = r
		}
		for i, r := range lineStr {
			s.SetContent(i, y, r, nil, lineNumberStyle)
		}

		// TODO: fix no cursor on empty line

		for x, c := range line {
			style := tcell.StyleDefault
			if activeLine == y {
				if t.Cursor.Position == x {
					style = cursorStyleCell
				} else {
					style = cursorStyle
				}
			}

			switch c {
			case '\n':
				s.SetContent(x+3, y, c, nil, cursorStyle)
			default:
				s.SetContent(x+3, y, c, nil, style)
			}

		}

		if activeLine == y && xx > len(line) {
			for x := len(line); x < xx; x++ {
				s.SetContent(x+3, y, ' ', nil, cursorStyle)
			}
		} else {
			for x := len(line); x < xx; x++ {
				s.SetContent(x+3, y, ' ', nil, tcell.StyleDefault)
			}

		}
	}
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
