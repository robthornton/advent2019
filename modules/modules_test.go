package modules_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/robthornton/advent2019/modules"
)

type mockReader []string

func (s mockReader) Read(buf []byte) (n int, err error) {
	str := strings.Join(s, "\n")
	return copy(buf, str), io.EOF
}

type badReader struct{}

func (b badReader) Read(buf []byte) (n int, err error) {
	return 0, errors.New("mock error")
}

func TestModuleMassFromBadReader(t *testing.T) {
	br := badReader{}
	modules, err := modules.FromReader(br)

	if err == nil {
		t.Fatal("expected non-nil error")
	}

	if len(modules) != 0 {
		t.Fatal("expected 0 modules, got: ", len(modules))
	}
}

func TestModuleMassesFromReader(t *testing.T) {
	mr := mockReader{"1", "2", "3", "4"}

	modules, err := modules.FromReader(mr)

	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	if len(modules) != len(mr) {
		t.Fatalf("expected %d modules, got %d", len(mr), len(modules))
	}
}

func TestModuleMassesFromReaderWithBadInput(t *testing.T) {
	mr := mockReader{"a"}

	modules, err := modules.FromReader(mr)

	if err == nil {
		t.Fatal("expected non-nil error")
	}

	if len(modules) != 0 {
		t.Fatal("expected 0 modules, got: ", len(modules))
	}
}

func TestModuleMassesFromReaderWithEndingNewline(t *testing.T) {
	mr := mockReader{"42\n"}

	modules, err := modules.FromReader(mr)

	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	if len(modules) != len(mr) {
		t.Fatalf("expected %d modules, got %d", len(mr), len(modules))
	}
}
