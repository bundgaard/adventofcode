package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode/utf8"
)

type requirement struct {
	Min      int
	Max      int
	Letter   rune
	Password string
}

type requirements []requirement

func parseRequirement(data string) (int, int, rune) {
	ins := strings.Split(data, " ")
	minmax, letter := ins[0], ins[1]
	run, _ := utf8.DecodeRuneInString(letter)
	arr := strings.Split(minmax, "-")
	min, _ := strconv.Atoi(arr[0])
	max, _ := strconv.Atoi(arr[1])
	return min, max, run
}

func verifyPassword(in requirement) bool {
	counter := 0
	for _, letter := range in.Password {
		// look for in.Letter
		if letter == in.Letter {
			counter++
		}

	}
	if counter >= in.Min && counter <= in.Max {
		return true
	}
	return false
}

func verifyPasswordPosition(in requirement) bool {
	// test if either position defined by min is occupied by in.Letter or if position max is occupied by in.Letter

	if byte(in.Letter) == in.Password[in.Min-1] && byte(in.Letter) == in.Password[in.Max-1] {
		return false
	} else if byte(in.Letter) == in.Password[in.Min-1] || byte(in.Letter) == in.Password[in.Max-1] {
		return true
	}

	return false
}
func main() {
	input, _ := ioutil.ReadFile("input")
	arr := strings.Split(string(input), "\n")

	requirements := make(requirements, 0)

	for _, val := range arr {

		ins := strings.FieldsFunc(val, func(r rune) bool {
			return r == ':'
		})

		min, max, run := parseRequirement(ins[0])
		password := strings.TrimSpace(ins[1])
		requirements = append(requirements, requirement{
			Min:      min,
			Max:      max,
			Letter:   run,
			Password: password})
	}

	// Part one
	counter := 0
	for _, req := range requirements {
		if verifyPassword(req) {
			counter++
		}
	} 	

	fmt.Println("Verified passwords (part one)", counter)

	// Part two

	counter = 0
	for _, req := range requirements {
		if verifyPasswordPosition(req) {
			counter++
		}
	}

	fmt.Println("Verified passwords (part two)", counter)

}
