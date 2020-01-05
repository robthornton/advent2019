package password_test

import (
	"testing"

	"github.com/robthornton/advent2019/password"
)

func TestValidElfPassword(t *testing.T) {
	tests := []struct {
		password string
		valid    bool
	}{
		{password: "111111", valid: true},
		{password: "223450", valid: false},
		{password: "123789", valid: false},
		{password: "123451", valid: false},
	}

	for _, test := range tests {
		if ok := password.ValidElfPassword(test.password); ok != test.valid {
			t.Errorf("expected password '%s' to be %+v but got %+v", test.password,
				test.valid, ok)
		}
	}
}

func TestValidElfPassword2(t *testing.T) {
	tests := []struct {
		password string
		valid    bool
	}{
		{password: "123444", valid: false},
		{password: "124444", valid: false},
		{password: "113334", valid: true},
		{password: "111334", valid: true},
		{password: "113345", valid: true},
		{password: "111122", valid: true},
		{password: "112233", valid: true},
		{password: "123445", valid: true},
		{password: "123456", valid: false},
	}

	for _, test := range tests {
		if ok := password.ValidElfPassword2(test.password); ok != test.valid {
			t.Errorf("expected password '%s' to be %+v but got %+v", test.password,
				test.valid, ok)
		}
	}
}
