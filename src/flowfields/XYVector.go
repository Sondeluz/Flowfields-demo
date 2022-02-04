package flowfields

type XYVector struct {
	X int
	Y int
}

func newXYVector(x int, y int) XYVector {
	return XYVector{x, y}
}

// Sets the default values for its fields
func (v *XYVector) setInvalid() {
	v.X = 0
	v.Y = 0
}

func (v *XYVector) isValid() bool {
	return (v.X != 0 || v.Y != 0)
}

// Returns true if the two vectors are going to collide (opposite direction on one of their axes, or both)
func (v1 XYVector) isColliding(v2 XYVector) bool {
	return -v1.X == v2.X && -v1.Y == v2.Y
}
