package main

import (
	"2021/util"
	"strings"
	"testing"
)

var doc = `3,4,3,1,2`

func TestNewLanternEvery6Day(t *testing.T) {

	// initial state
	lanternFish := util.ParseCommaSeparatedInputAsInts(strings.NewReader(doc))
	// day 01
	school := loadSchool(lanternFish)
	simulate(80, &school)

	t.Logf("Day 18: %d sum %d", school, sumSchool(&school))

}
