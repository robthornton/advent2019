package intcode_test

import (
	"strings"
	"testing"

	"github.com/robthornton/advent2019/intcode"
)

func TestNewIntcodeProgram(t *testing.T) {
	prog := intcode.NewProgram(strings.NewReader("99"))

	if len(prog.Errors()) != 0 {
		t.Fatalf("expected 0 errors, got: %d", len(prog.Errors()))
	}
}

func TestExecuteIntcodeHalt(t *testing.T) {
	prog := intcode.NewProgram(strings.NewReader("99"))
	prog.Run()

	if len(prog.Errors()) != 0 {
		for _, err := range prog.Errors() {
			t.Log(err)
		}
		t.Fatalf("expected 0 errors, got: %d", len(prog.Errors()))
	}
}

func TestExecuteIntcodeSet(t *testing.T) {
	prog := intcode.NewProgram(strings.NewReader("0"))
	prog.Set(0, 99)
	prog.Run()

	if len(prog.Errors()) != 0 {
		for _, err := range prog.Errors() {
			t.Log(err)
		}
		t.Fatalf("expected 0 errors, got: %d", len(prog.Errors()))
	}
}

func TestExecuteIntcodeUnknownCode(t *testing.T) {
	prog := intcode.NewProgram(strings.NewReader("0"))
	prog.Run()

	if len(prog.Errors()) != 1 {
		for _, err := range prog.Errors() {
			t.Log(err)
		}
		t.Fatalf("expected 1 error, got: %d", len(prog.Errors()))
	}
}

func TestExecuteIntcodeAdd(t *testing.T) {
	prog := intcode.NewProgram(strings.NewReader("1,5,6,5,99,10,32"))
	expected := "1,5,6,5,99,42,32"

	prog.Run()

	if len(prog.Errors()) != 0 {
		for _, err := range prog.Errors() {
			t.Log(err)
		}
		t.Fatalf("expected 0 errors, got: %d", len(prog.Errors()))
	}

	if prog.MemoryDump() != expected {
		t.Fatalf("expected %s, got %s", expected, prog.MemoryDump())
	}
}

func TestExecuteIntcodeMult(t *testing.T) {
	prog := intcode.NewProgram(strings.NewReader("2,5,6,6,99,10,42"))
	expected := "2,5,6,6,99,10,420"

	prog.Run()

	if len(prog.Errors()) != 0 {
		for _, err := range prog.Errors() {
			t.Log(err)
		}
		t.Fatalf("expected 0 errors, got: %d", len(prog.Errors()))
	}

	if prog.MemoryDump() != expected {
		t.Fatalf("expected %s, got %s", expected, prog.MemoryDump())
	}
}

func TestExecuteInputOutput(t *testing.T) {
	prog := intcode.NewProgram(strings.NewReader("3,0,4,0,99"))
	input := "42"

	prog.Input(input)
	prog.Run()

	if len(prog.Errors()) != 0 {
		for _, err := range prog.Errors() {
			t.Log(err)
		}
		t.Fatalf("expected 0 errors, got: %d", len(prog.Errors()))
	}

	if prog.Output() != input {
		t.Fatalf("expected %s, got %s", input, prog.Output())
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
		{src: "1101,100,-1,4,0", expected: "1101,100,-1,4,99"},
	}

	for _, test := range tests {
		prog := intcode.NewProgram(strings.NewReader(test.src))
		prog.Run()

		if len(prog.Errors()) != 0 {
			for _, err := range prog.Errors() {
				t.Log(err)
			}
			t.Fatalf("expected 0 errors, got: %d", len(prog.Errors()))
		}

		if prog.MemoryDump() != test.expected {
			t.Fatalf("expected %s, got %s", test.expected, prog.MemoryDump())
		}
	}
}

func TestRunImmediateMode(t *testing.T) {
	prog := intcode.NewProgram(strings.NewReader("1002,4,3,4,33"))
	prog.Run()

	if len(prog.Errors()) != 0 {
		for _, err := range prog.Errors() {
			t.Log(err)
		}
		t.Fatalf("expected 0 errors, got: %d", len(prog.Errors()))
	}
}

func TestRunEquality(t *testing.T) {
	tests := []struct {
		src      string
		input    string
		expected string
	}{
		{src: "3,9,8,9,10,9,4,9,99,-1,8", input: "8", expected: "1"}, // equal pos
		{src: "3,9,8,9,10,9,4,9,99,-1,8", input: "1", expected: "0"}, // not equal pos
		{src: "3,9,7,9,10,9,4,9,99,-1,8", input: "1", expected: "1"}, // less pos
		{src: "3,9,7,9,10,9,4,9,99,-1,8", input: "8", expected: "0"}, // not less pos
		{src: "3,3,1108,-1,8,3,4,3,99", input: "8", expected: "1"},   // equal imm
		{src: "3,3,1108,-1,8,3,4,3,99", input: "1", expected: "0"},   // not equal imm
		{src: "3,3,1107,-1,8,3,4,3,99", input: "1", expected: "1"},   // less imm
		{src: "3,3,1107,-1,8,3,4,3,99", input: "8", expected: "0"},   // not less imm
	}

	for _, test := range tests {
		prog := intcode.NewProgram(strings.NewReader(test.src))
		prog.Input(test.input)
		prog.Run()

		if len(prog.Errors()) != 0 {
			for _, err := range prog.Errors() {
				t.Log(err)
			}
			t.Fatalf("expected 0 errors, got: %d", len(prog.Errors()))
		}

		if prog.Output() != test.expected {
			t.Fatalf("expected %s, got %s", test.expected, prog.Output())
		}
	}
}

func TestRunJumps(t *testing.T) {
	tests := []struct {
		src      string
		input    string
		expected string
	}{
		{src: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", input: "1", expected: "1"}, // jump pos
		{src: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", input: "0", expected: "0"}, // not jump pos
		{src: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", input: "0", expected: "0"},      // jump imm
		{src: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", input: "1", expected: "1"},      // not jump imm
		{src: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: "7", expected: "999"},
		{src: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: "8", expected: "1000"},
		{src: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: "9", expected: "1001"},
	}

	for _, test := range tests {
		prog := intcode.NewProgram(strings.NewReader(test.src))
		prog.Input(test.input)
		prog.Run()

		if len(prog.Errors()) != 0 {
			for _, err := range prog.Errors() {
				t.Log(err)
			}
			t.Fatalf("expected 0 errors, got: %d", len(prog.Errors()))
		}

		if prog.Output() != test.expected {
			t.Fatalf("expected %s, got %s", test.expected, prog.Output())
		}
	}
}
