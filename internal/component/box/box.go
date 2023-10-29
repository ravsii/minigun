package box

import (
	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/component"
	"github.com/ravsii/minigun/internal/screen"
)

// RuneAt should implement drawing mechanism fox the box and return the rune
// at a given position.
// x and y are passed with removed offsets (margin, padding, etc), meaning
// even with padding and margins (0,0) will always be passed as a starting point.
//
// Box styles will be applied on Draw()
type RuneAt = func(x, y int) rune

type Box struct {
	x, y, w, h int

	fg, bg tcell.Color

	p      component.Padding
	border component.Border
}

func New(x, y, w, h int, opts ...BoxOption) *Box {
	b := Box{
		x:  x,
		y:  y,
		w:  w,
		h:  h,
		bg: tcell.ColorDefault,
		fg: tcell.ColorDefault,
	}

	for _, opt := range opts {
		opt(&b)
	}

	return &b
}

func (b *Box) Draw(runeAt RuneAt) {
	mask := b.border.Mask()
	borderLeft := mask&component.BorderLeft == component.BorderLeft
	borderTop := mask&component.BorderTop == component.BorderTop
	borderRight := mask&component.BorderRight == component.BorderRight
	borderBottom := mask&component.BorderBottom == component.BorderBottom

	fnTop := b.y + b.p.Top
	if borderTop {
		fnTop++
	}

	fnLeft := b.x + b.p.Left
	if borderLeft {
		fnLeft++
	}

	fullWidthEnd := fnLeft + b.w + b.p.Right
	if borderRight {
		fullWidthEnd++
	}
	fullHeightEnd := fnTop + b.h + b.p.Bottom
	if borderBottom {
		fullHeightEnd++
	}

	style := tcell.StyleDefault.Background(b.bg).Foreground(b.fg)

	// todo move padding/borders/etc draws inside them

	// Padding
	for y := b.y; y < fullHeightEnd; y++ {
		for x := b.x; x < fullWidthEnd; x++ {
			screen.SetEmptyStyle(x, y, style)
		}
	}

	// DrawFunc content
	for y := 0; y < b.h; y++ {
		for x := 0; x < b.w; x++ {
			r := runeAt(x, y)
			screen.SetRuneStyle(x+fnLeft, y+fnTop, r, style)
		}
	}

	borderRune := b.border.Rune()

	// Borders
	if borderTop {
		screen.FillLineFromTo(b.x, fullWidthEnd, b.y, borderRune, style)
	}
	if borderBottom {
		screen.FillLineFromTo(b.x, fullWidthEnd, fullHeightEnd, borderRune, style)
	}
	for y := b.y; y < fullHeightEnd; y++ {
		if borderLeft {
			screen.SetRuneStyle(b.x, y, borderRune, style)
		}
		if borderRight {
			screen.SetRuneStyle(fullWidthEnd-1, y, borderRune, style)
		}
	}
}
