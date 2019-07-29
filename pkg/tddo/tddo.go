package tddo

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

// Run genereting todo file
func Run() error {
	if ok := exists(FILENAME); ok {
		fmt.Printf("%s is already exists\n", FILENAME)
		return nil
	}
	fmt.Printf("? Generate %s in current directory? (Y/n)\n", FILENAME)
	if ok := confirm(os.Stdin); !ok {
		fmt.Println("No")
		return nil
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Printf("Generating %s in %s\n", FILENAME, dir)
	if err := generate(FILENAME, TEMPLATE); err != nil {
		return err
	}
	fmt.Printf("Successfully generated %s\n", FILENAME)
	return nil
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

func generate(name, text string) error {
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
