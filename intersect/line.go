package intersect

import "math"

// Line represents a line
type Line struct {
	x1, x2, y1, y2 int
}

// NewLine creates a new line.
func NewLine(x1, x2, y1, y2 int) Line {
	return Line{x1: x1, x2: x2, y1: y1, y2: y2}
}

// Length returns the length of a line.
func (l Line) Length() int {
	x := int(math.Abs(float64(l.x1 - l.x2)))
	y := int(math.Abs(float64(l.y1 - l.y2)))
	return x + y
}

// IntersectionsWithWire returns a list of points where Vector v has
// intersects with Vectors in wire w
func (l Line) IntersectionsWithWire(other Wire, distances *[]int) {
	for _, oline := range other.path {
		if point, ok := l.IntersectionWithLine(oline); ok {
			absX := int(math.Abs(float64(point.x)))
			absY := int(math.Abs(float64(point.y)))
			*distances = append(*distances, absX+absY)
		}
	}
}

// IntersectionWithLine returns a point and indicates whether an intersection
// was found. It will always return a point but the second argument will return
// false if a intersection was not found.
func (l Line) IntersectionWithLine(other Line) (Point, bool) {
	// problem does not necessitate parallel lines.
	if (l.x1 == l.x2 && other.x1 == other.x2) ||
		(l.y1 == l.y2 && other.y1 == other.y2) {
		return Point{}, false
	}

	if l.x1 < other.x1 && l.x1 < other.x2 && l.x2 < other.x1 && l.x2 < other.x2 {
		return Point{}, false
	}

	if l.x1 > other.x1 && l.x1 > other.x2 && l.x2 > other.x1 && l.x2 > other.x2 {
		return Point{}, false
	}

	if l.y1 < other.y1 && l.y1 < other.y2 && l.y2 < other.y1 && l.y2 < other.y2 {
		return Point{}, false
	}

	if l.y1 > other.y1 && l.y1 > other.y2 && l.y2 > other.y1 && l.y2 > other.y2 {
		return Point{}, false
	}

	var x, y int
	if l.x1 == l.x2 {
		x = l.x1
	}

	if other.x1 == other.x2 {
		x = other.x1
	}

	if l.y1 == l.y2 {
		y = l.y1
	}

	if other.y1 == other.y2 {
		y = other.y1
	}

	// a match of origin (0, 0) is prohibited
	if x+y == 0 {
		return Point{}, false
	}

	return Point{x, y}, true
}
