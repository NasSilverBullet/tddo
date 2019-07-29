package tddo

import (
	"bytes"
	"testing"
)

func TestExist(t *testing.T) {
	existsTests := []struct {
		in  string
		out bool
	}{
		{"tddo.go", true},
		{"hoge.go", false},
	}
	for _, test := range existsTests {
		td := &Tddo{Name: test.in}
		t.Run(test.in, func(t *testing.T) {
			if expected, actual := test.out, td.Exists(); expected != actual {
				t.Errorf("td.Exists wont %t but got %t", expected, actual)
			}

		})
	}

}

func TestConfirm(t *testing.T) {
	confirmTests := []struct {
		in  string
		out bool
	}{
		{"Y", true},
		{"n", false},
		{"other", false},
	}
	for _, test := range confirmTests {
		t.Run(test.in, func(t *testing.T) {
			td := &Tddo{Reader: bytes.NewBufferString(test.in)}
			if expected, actual := test.out, td.Confirm(); expected != actual {
				t.Errorf("td.confirm() wont %t but got %t", expected, actual)
			}

		})
	}
}
