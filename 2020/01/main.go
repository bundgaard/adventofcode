package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func readInts(fp string, sep string) []int {
	input, _ := ioutil.ReadFile(fp)
	arr := make([]int, 0)
	_ = arr
	for _, val := range strings.Split(string(input), sep) {
		n, _ := strconv.Atoi(val)
		arr = append(arr, n)
	}
	return arr
}

func partOne(ints []int) int {

	for i := range ints {
		for j := range ints {
			a := ints[i]
			b := ints[j]
			if i == j {
				continue
			}

			if a+b == 2020 {
				result := a * b
				fmt.Println(result)
				return result
			}
		}
	}
	return 0
}

func partTwo(ints []int) int {

	for i := range ints {
		for j := range ints {
			for k := range ints {
				a := ints[i]
				b := ints[j]
				c := ints[k]
				if i == j && i == k {
					continue
				}

				if a+b+c == 2020 {
					result := a * b * c
					fmt.Println(result)
					return result
				}
			}

		}
	}
	return 0
}
func main() {
	ints := readInts("input.txt", "\n")
	sort.Ints(ints)
	fmt.Println("part one", partOne(ints))
	/// part two
	fmt.Println("part two", partTwo(ints))

}
