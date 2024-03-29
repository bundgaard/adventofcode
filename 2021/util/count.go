package util

import (
	"math"
)

type Value struct {
	Bit int
}

func CountBitsInIntSlice(arr []Value) map[int]int {
	result := make(map[int]int)
	for _, value := range arr {
		result[value.Bit]++
	}
	return result
}

func MostCommonBitIntSlice(arr []Value) int {
	result := CountBitsInIntSlice(arr)
	currentMax := math.MinInt
	maxBit := 0
	for bit, count := range result {
		if count > currentMax {
			currentMax = count
			maxBit = bit
		}
	}
	return maxBit
}

func LeastCommonBitIntSlice(arr []Value) int {
	result := CountBitsInIntSlice(arr)
	currentMin := math.MaxInt
	minBit := 0
	for bit, count := range result {
		if count < currentMin {
			currentMin = count
			minBit = bit
		}
	}
	return minBit
}
