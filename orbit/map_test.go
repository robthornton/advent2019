package orbit_test

import (
	"strings"
	"testing"

	"github.com/robthornton/advent2019/orbit"
)

func TestNewOrbit(t *testing.T) {
	com := orbit.NewMap()

	if com.Count() != 0 {
		t.Fatalf("expected %d got %d", 0, com.Count())
	}
}

func TestInsertOneOrbit(t *testing.T) {
	com := orbit.NewMap()
	com.Insert("AAA", "BBB")

	expected := 2

	if com.Count() != expected {
		t.Fatalf("expected %d got %d", expected, com.Count())
	}
}

func TestBuildTree(t *testing.T) {
	orbitMap :=
		`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

	com := orbit.NewFromString(orbitMap)
	expected := 12

	if com.Count() != expected {
		t.Fatalf("expected count of %d got %d", expected, com.Count())
	}
}

func TestChecksum(t *testing.T) {
	orbitMap := `
COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
`

	tree := orbit.NewFromReader(strings.NewReader(orbitMap))

	expected := 42
	checksum := tree.Checksum()

	if checksum != expected {
		t.Fatalf("expected count of %d got %d", expected, checksum)
	}
}

func TestFewestTransfers(t *testing.T) {
	orbitMap := `
COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN
`

	m := orbit.NewFromReader(strings.NewReader(orbitMap))
	expected := 4
	transfers := m.FewestTransfers("YOU", "SAN")

	if transfers != expected {
		t.Fatalf("expected count of %d got %d", expected, transfers)
	}
}
