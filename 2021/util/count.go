package util

import "math"

func CountBitsInIntSlice(arr []int) map[int]int {
	result := make(map[int]int)
	for _, bit := range arr {
		result[bit]++
	}
	return result
}

func MostCommonBitIntSlice(arr []int) int {
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

func LeastCommonBitIntSlice(arr []int) int {
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
