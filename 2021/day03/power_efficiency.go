package main

import (
	"2021/util"
	"log"
	"sort"
	"strconv"
)

func findLeastMostBits(bits map[int][]util.Value) ([]int, []int) {
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

func calculatePowerEfficiencyη(bits map[int][]util.Value) (int64, int64) {
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
