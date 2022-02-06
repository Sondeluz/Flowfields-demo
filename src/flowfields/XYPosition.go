package flowfields

type XYPosition struct {
	X int
	Y int
}

func (v *XYPosition) setInvalid() {
	v.X = -1
	v.Y = -1
}

// Returns true if the position is valid (non-negative on both axes)
func (p *XYPosition) isValid() bool {
	return (p.X != -1 && p.Y != -1)
}

func (p1 XYPosition) equals(p2 XYPosition) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

// Return the resulting position from advancing it with the given direction
func (pos *XYPosition) advance(v XYVector) (newPos XYPosition) {
	newPos.X = pos.X + v.X
	newPos.Y = pos.Y + v.Y
	return newPos
}
