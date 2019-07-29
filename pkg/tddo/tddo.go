package tddo

import (
	"bufio"
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

// Tddo has file's data
type Tddo struct {
	Name   string
	Text   string
	Reader io.Reader
}

// New is Tddo constructor
func New() *Tddo {
	return &Tddo{
		Name:   FILENAME,
		Text:   TEMPLATE,
		Reader: os.Stdin,
	}
}

// Exists check file exists
func (td *Tddo) Exists() bool {
	_, err := os.Stat(td.Name)
	return !os.IsNotExist(err)
}

// Confirm inquire to user
func (td *Tddo) Confirm() bool {
	sc := bufio.NewScanner(td.Reader)
	sc.Scan()
	if sc.Text() != "Y" {
		return false
	}
	return true
}

// Generate files with given name and text contents
func (td *Tddo) Generate() error {
	f, err := os.OpenFile(td.Name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(td.Text))
	if err != nil {
		return err
	}
	return nil
}
