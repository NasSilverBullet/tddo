package main

import "testing"

func TestExist(t *testing.T) {
	fn := "README.md"
	if expected, actual := exists(fn), true; expected != actual {
		t.Errorf("exists(%s) wont %t but got %t", fn, expected, actual)
	}
}
