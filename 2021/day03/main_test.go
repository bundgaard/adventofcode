package main

import (
	"2021/util"
	"log"
	"strings"
	"testing"
)

var doc = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestPart01(t *testing.T) {

	bits := run(util.ParseInput(strings.NewReader(doc)))

	ε, γ := calculatePowerEfficiencyη(bits)
	t.Log("ε =", ε)
	t.Log("γ =", γ)

	t.Log("ε * γ = ", ε*γ)
}
func Count(in map[int]int) map[int]int {
	out := make(map[int]int)

	for k, v := range in {

		log.Printf("count %d %d", k, v)
	}
	return out

}

type Index int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func TestPart02(t *testing.T) {
	lines := util.ParseInput(strings.NewReader(doc))
	mapIndexAndBinary := run(lines)
	t.Logf("bits %v", mapIndexAndBinary)
	// Loop over all numbers and find which has count > most significant
	for pos, bits := range mapIndexAndBinary {
		t.Logf("position %d value %v", pos, bits)
		freqMap := util.CountBitsInIntSlice(bits)
		if freqMap[0] == freqMap[1] {

		}
		// mostSignificant := max(freqMap[0], freqMap[1])
	}

	t.Logf("lines %v", lines)

}
