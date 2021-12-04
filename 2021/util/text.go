package util

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func ReverseInts(arr []int) []int {
	result := make([]int, 0)
	for i := len(arr) - 1; i > -1; i-- {
		result = append(result, arr[i])

	}
	return result
}

func IntToString(arr []int) string {
	var out strings.Builder
	for _, e := range arr {
		out.WriteString(fmt.Sprint(e))
	}
	return out.String()
}

func StringToInts(aString string) []int {
	result := make([]int, 0)
	for _, chr := range aString {
		result = append(result, int(chr-'0'))
	}
	return result
}
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
