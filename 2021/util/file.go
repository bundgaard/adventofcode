package util

import (
	"log"
	"os"
)

func OpenFile(filename string) *os.File {
	fd, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return fd
}
