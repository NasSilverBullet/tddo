package main

import (
	"fmt"
	"os"
)

// FILENAME is target todo file.
const FILENAME = "tddo.md"

func run() {
	if ok := exists(FILENAME); ok {
		fmt.Printf("%s is already exists", FILENAME)
	}
}

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
