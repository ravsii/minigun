package tab

type Cursor struct {
	Line     int
	Position int
	// prevPosition is used to keep horizontal positing when moving up and down
	// between the lines of different width.
	prevPosition int
}
