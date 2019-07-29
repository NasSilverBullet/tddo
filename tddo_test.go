package main

import (
	"bytes"
	"testing"
)

func TestExist(t *testing.T) {
	existsTests := []struct {
		in  string
		out bool
	}{
		{"README.md", true},
		{"hoge.md", false},
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
