package util

import (
	"bufio"
	"io"
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

func ParseInput(r io.Reader) []string {
	br := bufio.NewScanner(r)
	result := make([]string, 0)
	for br.Scan() {
		line := br.Text()
		result = append(result, line)
	}
	return result
}
