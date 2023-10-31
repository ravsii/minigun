package tabs

type Cursor struct {
	Line     int
	Position int
	// PrevPosition is used to keep horizontal positing when moving up and down
	// between the lines of different width.
	PrevPosition int
}
