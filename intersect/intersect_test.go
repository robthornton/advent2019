package intersect_test

import (
	"testing"

	"github.com/robthornton/advent2019/intersect"
)

func TestLineSegmentIntersects(t *testing.T) {
	tests := []struct {
		a, b       intersect.Line
		intersects bool
		dist       int
	}{
		// match
		{
			a:          intersect.NewLine(0, 2, 1, 1),
			b:          intersect.NewLine(1, 1, 0, 2),
			intersects: true,
			dist:       2,
		},
		// right
		{
			a:          intersect.NewLine(0, 1, 2, 1),
			b:          intersect.NewLine(3, 0, 3, 2),
			intersects: false,
			dist:       0,
		},
		// left
		{
			a:          intersect.NewLine(1, 1, 3, 1),
			b:          intersect.NewLine(0, 0, 0, 2),
			intersects: false,
			dist:       0,
		},
		// above
		{
			a:          intersect.NewLine(0, 3, 2, 3),
			b:          intersect.NewLine(1, 0, 1, 2),
			intersects: false,
			dist:       0,
		},
		// below
		{
			a:          intersect.NewLine(0, 1, 2, 1),
			b:          intersect.NewLine(1, 3, 1, 5),
			intersects: false,
			dist:       0,
		},
		// parallel horizontal
		{
			a:          intersect.NewLine(0, 0, 2, 0),
			b:          intersect.NewLine(0, 1, 2, 1),
			intersects: false,
			dist:       0,
		},
		// parallel vertical
		{
			a:          intersect.NewLine(0, 0, 0, 2),
			b:          intersect.NewLine(1, 0, 1, 2),
			intersects: false,
			dist:       0,
		},
	}

	for i, test := range tests {
		dist, ok := test.a.IntersectionWithLine(test.b)
		if ok != test.intersects {
			t.Errorf("(%d) expected match to be %v but got %v", i, test.intersects, ok)
		}

		if dist.X()+dist.Y() != test.dist {
			t.Errorf("(%d) expected match to be %d but got %d", i, test.dist, dist)
		}
	}
}

func TestNearestWireIntersect(t *testing.T) {
	tests := []struct {
		paths    [2]string
		expected int
	}{
		{paths: [2]string{
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
		}, expected: 159},
		{paths: [2]string{
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
		}, expected: 135},
	}

	for i, test := range tests {
		wireA := intersect.NewWireFromPath(test.paths[0])
		wireB := intersect.NewWireFromPath(test.paths[1])
		dist, ok := wireA.ClosestIntersectionToOrigin(wireB)

		if !ok {
			t.Error("no match found")
		}

		if dist != test.expected {
			t.Errorf("%d: expected distance of %d but got %d", i, test.expected, dist)
		}
	}
}

func TestShortestDistanceIntersect(t *testing.T) {
	tests := []struct {
		paths    [2]string
		expected int
	}{
		{paths: [2]string{
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
		}, expected: 610},
		{paths: [2]string{
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
		}, expected: 410},
	}

	for i, test := range tests {
		wireA := intersect.NewWireFromPath(test.paths[0])
		wireB := intersect.NewWireFromPath(test.paths[1])
		dist := wireA.ShortestIntersect(wireB)

		if dist != test.expected {
			t.Errorf("%d: expected distance of %d but got %d", i, test.expected, dist)
		}
	}
}
