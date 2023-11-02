package tabs

import (
	"fmt"
	"math"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/components"
	"github.com/ravsii/minigun/internal/components/box"
)

const vbar = '┆'

func (t *Tab) lineNumbersWidth() int {
	// +1 because it starts at 0
	lines := len(t.lines) + 1
	// neededWidth stores the amount of runes
	// needed to represent the amount of lines.
	neededWidth := int(math.Ceil(math.Log10(float64(lines))))
	return neededWidth
}

func (t *Tab) drawLineNumbers(fromLine int) {
	border := components.NewBorder(vbar, components.BorderRight)

	width := t.lineNumbersWidth()

	lineBox := box.New(t.xOffset, t.yOffset, width, t.h,
		box.WithBorder(border),
		box.WithPadding(components.Padding{Left: 1}),
		box.WithTextColor(tcell.ColorWhite.TrueColor()),
	)

	lines := make([]string, t.h)
	for y := 0; y < t.h; y++ {
		line := fromLine + y
		if line < fromLine+t.h {
			lines[y] = fmt.Sprintf("%*d", width, line+1)
		} else {
			lines[y] = fmt.Sprintf("%*c", width, '~')
		}
	}

	lineBox.Draw(func(x, y int) rune {
		runes := []rune(lines[y])
		return runes[x]
	})
}
