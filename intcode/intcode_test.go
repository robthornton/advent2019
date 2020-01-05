package main

import (
	"strings"
	"testing"
)

func TestNewIntcodeProgram(t *testing.T) {
	prog := newIntcodeProgram(strings.NewReader("99"))

	if len(prog.errors) != 0 {
		t.Fatalf("expected 0 errors, got: %d", len(prog.errors))
	}
}

func TestExecuteIntcodeHalt(t *testing.T) {
	prog := newIntcodeProgram(strings.NewReader("99"))
	prog.run()

	if len(prog.errors) != 0 {
		for _, err := range prog.errors {
			t.Log(err)
		}
		t.Fatalf("expected 0 errors, got: %d", len(prog.errors))
	}
}

func TestExecuteIntcodeSet(t *testing.T) {
	prog := newIntcodeProgram(strings.NewReader("0"))
	prog.set(0, 99)
	prog.run()

	if len(prog.errors) != 0 {
		for _, err := range prog.errors {
			t.Log(err)
		}
		t.Fatalf("expected 0 errors, got: %d", len(prog.errors))
	}
}

func TestExecuteIntcodeUnknownCode(t *testing.T) {
	prog := newIntcodeProgram(strings.NewReader("0"))
	prog.run()

	if len(prog.errors) != 1 {
		for _, err := range prog.errors {
			t.Log(err)
		}
		t.Fatalf("expected 1 error, got: %d", len(prog.errors))
	}
}

func TestExecuteIntcodeAdd(t *testing.T) {
	prog := newIntcodeProgram(strings.NewReader("1,5,6,5,99,10,32"))
	expected := "1,5,6,5,99,42,32"

	prog.run()

	if len(prog.errors) != 0 {
		for _, err := range prog.errors {
			t.Log(err)
		}
		t.Fatalf("expected 0 errors, got: %d", len(prog.errors))
	}

	if prog.out() != expected {
		t.Fatalf("expected %s, got %s", expected, prog.out())
	}
}

func TestExecuteIntcodeMult(t *testing.T) {
	prog := newIntcodeProgram(strings.NewReader("2,5,6,6,99,10,42"))
	expected := "2,5,6,6,99,10,420"

	prog.run()

	if len(prog.errors) != 0 {
		for _, err := range prog.errors {
			t.Log(err)
		}
		t.Fatalf("expected 0 errors, got: %d", len(prog.errors))
	}

	if prog.out() != expected {
		t.Fatalf("expected %s, got %s", expected, prog.out())
	}
}

func TestExecuteIntcodePrograms(t *testing.T) {
	tests := []struct {
		src      string
		expected string
	}{
		{src: "1,0,0,0,99", expected: "2,0,0,0,99"},
		{src: "2,3,0,3,99", expected: "2,3,0,6,99"},
		{src: "2,4,4,5,99,0", expected: "2,4,4,5,99,9801"},
		{src: "1,1,1,4,99,5,6,0,99", expected: "30,1,1,4,2,5,6,0,99"},
	}

	for _, test := range tests {
		prog := newIntcodeProgram(strings.NewReader(test.src))
		prog.run()

		if len(prog.errors) != 0 {
			for _, err := range prog.errors {
				t.Log(err)
			}
			t.Fatalf("expected 0 errors, got: %d", len(prog.errors))
		}

		if prog.out() != test.expected {
			t.Fatalf("expected %s, got %s", test.expected, prog.out())
		}
	}
}
