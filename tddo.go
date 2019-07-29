package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	// FILENAME is target todo file.
	FILENAME = "tddo.md"

	// TEMPLATE is todo file template.
	TEMPLATE = `# TODO for test-driven development

- [x] Create TODO
- [ ]
- [ ]
- [ ]
- [ ]
`
)

func run() {
	if ok := exists(FILENAME); !ok {
		fmt.Printf("%s is already exists", FILENAME)
		return
	}
	confirm(os.Stdin)
}

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func confirm(r io.Reader) bool {
	sc := bufio.NewScanner(r)
	sc.Scan()
	if sc.Text() != "Y" {
		return false
	}
	return true
}

func create(name, text string) error {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0755)
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(text))
	if err != nil {
		return err
	}
	return nil
}
