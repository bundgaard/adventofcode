package util

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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

func ParseCommaSeparatedInputAsInts(r io.Reader) []int {
	br := bufio.NewScanner(r)
	result := make([]int, 0)
	for br.Scan() {
		line := br.Text()
		digits := strings.Split(line, ",")
		for _, digit := range digits {
			n, err := strconv.Atoi(digit)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, n)
		}
	}
	return result
}
