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

	p component.Padding
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
	fnTop := b.y + b.p.Top
	// fnBottom := b.y + b.h
	fnLeft := b.x + b.p.Left
	// fnRight := b.x + b.w

	fullWidthEnd := b.x + b.w + b.p.Left + b.p.Right
	fullHeightEnd := b.y + b.h + b.p.Top + b.p.Bottom

	style := tcell.StyleDefault.Background(b.bg).Foreground(b.fg)

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

	screen.Show()
}
