package main

import "2021/util"

func CO2ScrubberRating(arr []string, pos int) []string {
	if len(arr) == 1 {
		return arr
	}
	l := len(arr[0])
	for i := 0; i < l; i++ {
		arr = bitcriteriaCO2ScrubberRating(arr, i)
	}
	return arr
}

func oxygenGeneratorRating(arr []string, pos int) []string {
	if len(arr) == 1 {
		return arr
	}
	l := len(arr[0])
	for i := 0; i < l; i++ {
		arr = bitCriteriaOxygen(arr, i)
	}
	return arr
}

func bitCriteriaOxygen(arr []string, pos int) []string {
	counter := make(map[byte]int)

	for _, elm := range arr {
		counter[elm[pos]]++
	}
	if counter[49] == counter[48] {
		for _, line := range arr {
			if line[pos] == 49 {
				return []string{line}
			}
		}
	}
	k := util.MaxMap(counter)
	result := make([]string, 0)
	for _, line := range arr {
		if line[pos] == k {
			result = append(result, line)
		}
	}
	return result
}

func bitcriteriaCO2ScrubberRating(arr []string, pos int) []string {
	if len(arr) == 1 {
		return arr
	}
	counter := make(map[byte]int)
	for _, elm := range arr {
		counter[elm[pos]]++
	}
	if counter[49] == counter[48] {
		for _, elm := range arr {
			if elm[pos] == 48 {
				return []string{elm}
			}
		}
	}
	k := util.MinMap(counter)

	result := make([]string, 0)
	for _, elm := range arr {
		if elm[pos] == k {
			result = append(result, elm)
		}
	}
	return result
}
