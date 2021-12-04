package main

import (
	"2021/util"
	"fmt"
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
		result = append(result, sum)
	}
	return result
}

func main() {

	fd := util.OpenFile("day01.txt")
	defer fd.Close()
	measurements := util.ParseInputAsInts(fd)
	fmt.Println("Part 01")
	fmt.Printf("increments %-4d\n", countIncrement(measurements))
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Part 02")
	fmt.Printf("Sliding increments %d\n", countIncrement(countSlidingWindow(measurements, 3)))

}
