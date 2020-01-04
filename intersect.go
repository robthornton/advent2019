package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Point is a set of cartesian coordinates
type Point struct {
	x, y int
}

// Line represents a line
type Line struct {
	x1, x2, y1, y2 int
}

// Wire contains the path the wire follows
type Wire struct {
	path []Line
}

// NewWireFromPath creates a new Wire object from a command string
func NewWireFromPath(path string) Wire {
	commands := strings.Split(path, ",")
	current := Point{}
	lines := make([]Line, len(commands))

	for i, command := range commands {
		next := Point{x: current.x, y: current.y}
		dist, _ := strconv.ParseInt(command[1:], 10, 64)

		switch command[0] {
		case 'D':
			next.y -= int(dist)
		case 'L':
			next.x -= int(dist)
		case 'R':
			next.x += int(dist)
		case 'U':
			next.y += int(dist)
		}

		l := Line{
			x1: current.x,
			x2: next.x,
			y1: current.y,
			y2: next.y,
		}

		lines[i] = l
		current = next
	}

	return Wire{path: lines}
}

// ClosestIntersectionToOrigin returns the distance to the closest intersection
// and whether an intersect was actually found.
func (w Wire) ClosestIntersectionToOrigin(other Wire) (int, bool) {
	distances := w.IntersectionsWithWire(other)

	sort.IntSlice(distances).Sort()

	return distances[0], true
}

// ShortestIntersect calculates the shortest distance to the first intersection
// of two wires
func (w Wire) ShortestIntersect(other Wire) int {
	distances := make([]int, 0)
	distanceA := 0

	for _, l1 := range w.path {
		distanceB := 0

		for _, l2 := range other.path {
			if p, ok := l1.IntersectionWithLine(l2); ok {
				// calculate the segments to the intersection and add then into the total
				// distance travelled on each wire
				l3 := Line{x1: l1.x1, x2: p.x, y1: l1.y1, y2: p.y}
				l4 := Line{x1: l2.x1, x2: p.x, y1: l2.y1, y2: p.y}
				distances = append(distances, distanceA+distanceB+l3.Length()+l4.Length())
			}
			distanceB += l2.Length()
		}

		distanceA += l1.Length()
	}

	sort.IntSlice(distances).Sort()

	fmt.Println(distances)

	return distances[0]
}

func (l Line) Length() int {
	x := int(math.Abs(float64(l.x1 - l.x2)))
	y := int(math.Abs(float64(l.y1 - l.y2)))
	return x + y
}

// IntersectionsWithWire returns a list of points where Vectors in wire w
// intersect with Vectors in wire other
func (w Wire) IntersectionsWithWire(other Wire) []int {
	distances := make([]int, 0)

	for _, l := range w.path {
		l.IntersectionsWithWire(other, &distances)
	}

	return distances
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
