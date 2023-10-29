package box

import (
	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/component"
)

type BoxOption func(*Box)

// WithPadding adds padding to the box item.
func WithPadding(p component.Padding) BoxOption {
	return func(b *Box) {
		b.p = p
	}
}

// WithTextColor sets text color.
func WithTextColor(c tcell.Color) BoxOption {
	return func(b *Box) {
		b.fg = c
	}
}

// WithBackground sets background color.
func WithBackground(c tcell.Color) BoxOption {
	return func(b *Box) {
		b.bg = c
	}
}
