package orbit

import (
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

// Map is an Orbit map.
type Map map[string][]string

// NewMap creates a new Orbit map.
func NewMap() Map {
	return make(Map)
}

// NewFromReader builds a tree from the given orbit string
func NewFromReader(r io.Reader) Map {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil
	}

	return NewFromString(string(buf))
}

// NewFromString builds a tree from the given orbit string
func NewFromString(s string) Map {
	s = strings.Trim(s, "\n")
	orbits := strings.Split(s, "\n")
	m := NewMap()

	for _, orbit := range orbits {
		IDs := strings.Split(orbit, ")")

		if len(IDs) != 2 {
			continue
		}

		m.Insert(IDs[0], IDs[1])
	}

	return m
}

// Checksum count of direct and indirect orbits
func (m Map) Checksum() int {
	checksum := 0

	// foreach parent, count until a match not found.
	for _, children := range m {
		for _, child := range children {
			checksum += m.Depth(child)
		}
	}

	return checksum
}

// Depth returns the depth of the given parent.
func (m Map) Depth(ID string) int {
	depth := 0

	for parent, children := range m {
		for _, child := range children {
			if child == ID {
				return 1 + m.Depth(parent)
			}
		}
	}

	return depth
}

// Count returns the total number of elements in the map
func (m Map) Count() int {
	list := make(map[string]bool)
	for parent, children := range m {
		list[parent] = true

		for _, child := range children {
			list[child] = true
		}
	}

	return len(list)
}

// Insert a new parent and child relationship
func (m Map) Insert(parent, child string) {
	list, ok := m[parent]
	if !ok {
		list = make([]string, 0)
	}
	list = append(list, child)
	m[parent] = list
}

// FewestTransfers calculates the fewest orbital transfers to get from YOU
// to SAN
func (m Map) FewestTransfers(from, to string) int {
	pathFrom := make([]string, 0)
	pathTo := make([]string, 0)

	path := m.GetParent(from)
	for path != "" {
		pathFrom = append(pathFrom, path)
		path = m.GetParent(path)
	}

	path = m.GetParent(to)
	for path != "" {
		pathTo = append(pathTo, path)
		path = m.GetParent(path)
	}

	distances := make(sort.IntSlice, 0)
	for i, from := range pathFrom {
		for j, to := range pathTo {
			if to == from {
				distances = append(distances, i+j)
			}
		}
	}

	distances.Sort()

	return distances[0]
}

// GetParent of child ID
func (m Map) GetParent(ID string) string {
	for parent, children := range m {
		for _, child := range children {
			if child == ID {
				return parent
			}
		}
	}

	return ""
}
