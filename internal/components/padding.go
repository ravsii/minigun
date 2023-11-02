package components

type Padding struct {
	Top, Bottom, Left, Right int
}

func (p *Padding) SumX() int {
	return p.Left + p.Right
}

func (p *Padding) SumY() int {
	return p.Top + p.Bottom
}
