package main

import (
	"2021/util"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Value struct {
	Number  string
	Ordinal int
}

func run(arr []string) map[int][]int {
	bits := make(map[int][]int)

	for _, line := range arr {
		for position, bit := range line {
			bits[position] = append(bits[position], int(bit)-'0')
		}
	}
	return bits
}
func findLeastMostBits(bits map[int][]int) ([]int, []int) {
	sortedKey := make([]int, 0, len(bits))
	for k := range bits {
		sortedKey = append(sortedKey, k)
	}
	sort.Ints(sortedKey)
	var mostBits []int
	var leastBits []int
	for _, k := range sortedKey {
		mostBit := util.MostCommonBitIntSlice(bits[k])
		leastBit := util.LeastCommonBitIntSlice(bits[k])
		mostBits = append(mostBits, mostBit)
		leastBits = append(leastBits, leastBit)
	}
	return leastBits, mostBits
}

func calculatePowerEfficiencyη(bits map[int][]int) (int64, int64) {
	leastBits, mostBits := findLeastMostBits(bits)
	epsilon := util.IntToString(mostBits)
	ε, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	gamma := util.IntToString(leastBits)
	γ, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return ε, γ
}

func calculateLifeSupportRating(bits map[int][]int) {
	/* var (
		oxygenGeneratorRating int
		co2ScrubberRating     int
	) */

}

// bit criteria for oxygen
func oxygenGeneratorRating(arr [][]int) {
	// Determine the most common value of bit in the current position and keep only numbers with that bit
	// convert into map with mostCommonBits

}
func main() {

	fd := util.OpenFile("day03.txt")
	defer fd.Close()
	lines := util.ParseInput(fd)
	bits := run(lines)
	epsilon, gamma := calculatePowerEfficiencyη(bits)
	log.Println("Part 01")
	log.Printf("%d * %d = %d\n", epsilon, gamma, epsilon*gamma)
	log.Println(strings.Repeat("=", 80))

	// Convert list of string to list of list of ints

}
