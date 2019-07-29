package main

import (
	"log"

	"github.com/NasSilverBullet/tddo/pkg/tddo"
)

func main() {
	err := tddo.Run()
	if err != nil {
		log.Fatal(err)
	}
}
