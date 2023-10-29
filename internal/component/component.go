package component

// probably should be removed

type Component interface {
	// Draw should redraw the component.
	// It guarantees that component will be redrawn and shown on the screen.
	Draw()
}
