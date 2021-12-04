package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func countIncrement(ints []int) int {
	var count int
	last := ints[0]
	for i := 1; i < len(ints); i++ {
		current := ints[i]
		if current > last {
			count++
		}
		last = current
	}
	return count
}

func countSlidingWindow(arr []int, period int) []int {
	result := make([]int, 0)

	for i := 0; i < len(arr)-2; i++ {
		sum := 0
		for j := i; j < i+period; j++ {
			sum += arr[j]
		}
		log.Printf("Sum %d", sum)
		result = append(result, sum)
	}
	return result
}

func parseInputAsInts(r io.Reader) []int {
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

func main() {

	fd, err := os.Open("day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	measurements := parseInputAsInts(fd)
	fmt.Printf("increments %-4d\n", countIncrement(measurements))

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("Sliding increments %d\n", countIncrement(countSlidingWindow(measurements, 3)))

}
