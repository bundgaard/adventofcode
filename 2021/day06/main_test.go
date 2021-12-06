package main

import (
	"2021/util"
	"strings"
	"testing"
)

var doc = `3,4,3,1,2`

func TestNewLanternEvery6Day(t *testing.T) {

	// initial state
	lanternSchool := util.ParseCommaSeparatedInputAsInts(strings.NewReader(doc))
	// day 01
	lanternSchool = simulate(80, lanternSchool)
	t.Logf("Day 18: %d", len(lanternSchool))

}
