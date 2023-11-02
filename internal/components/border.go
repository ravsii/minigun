package components

type BorderPosition int

const (
	BorderTop BorderPosition = 1 << iota
	BorderLeft
	BorderRight
	BorderBottom
	BorderAll
)

type Border struct {
	r    rune
	mask BorderPosition
}

// NewBorder creates a new border element with a given rune.
// borderMask accept a mask-like arguments for the border positions.
// For example:
//
//	NewBorder('|', BorderLeft | BorderRight) // Apply left & right borders
//	NewBorder('|', BorderAll) // Apply all borders
func NewBorder(borderRune rune, borderMask BorderPosition) Border {
	return Border{
		r:    borderRune,
		mask: borderMask,
	}
}

func (b *Border) Rune() rune {
	return b.r
}

func (b *Border) Mask() BorderPosition {
	return b.mask
}
