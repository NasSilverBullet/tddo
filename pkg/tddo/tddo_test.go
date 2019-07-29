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
		t.Run(test.in, func(t *testing.T) {
			if expected, actual := test.out, exists(test.in); expected != actual {
				t.Errorf("esists(%s) wont %t but got %t", test.in, expected, actual)
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
			if expected, actual := test.out, confirm(bytes.NewBufferString(test.in)); expected != actual {
				t.Errorf("confirm(%s) wont %t but got %t", test.in, expected, actual)
			}

		})
	}
}
