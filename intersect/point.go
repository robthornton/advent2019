package intersect

// Point is a set of cartesian coordinates
type Point struct {
	x, y int
}

// NewPoint returns a new point object.
func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

// X returns the x coordinate of the point.
func (p Point) X() int {
	return p.x
}

// Y returns the y coordinate of the point.
func (p Point) Y() int {
	return p.y
}
