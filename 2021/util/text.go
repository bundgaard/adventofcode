package util

import (
	"bufio"
	"io"
	"log"
	"strconv"
)

func ParseInputAsInts(r io.Reader) []int {
	result := make([]int, 0)
	br := bufio.NewScanner(r)
	for br.Scan() {

		n, err := strconv.Atoi(br.Text())
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, n)
	}
	return result
}
