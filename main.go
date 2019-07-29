package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NasSilverBullet/tddo/pkg/tddo"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	td := tddo.New()
	if ok := td.Exists(); ok {
		fmt.Printf("%s is already exists\n", td.Name)
		return nil
	}
	fmt.Printf("? Generate %s in current directory? (Y/n)\n", td.Name)
	if ok := td.Confirm(); !ok {
		fmt.Println("No")
		return nil
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Printf("Generating %s in %s\n", td.Name, dir)
	if err := td.Generate(); err != nil {
		return err
	}
	fmt.Printf("Successfully generated %s\n", td.Name)
	return nil
}
