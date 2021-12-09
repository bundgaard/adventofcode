package day07

import (
	"2021/util"
	"strings"
	"testing"
)

var doc = `16,1,2,0,4,2,7,1,2,14`

func TestPart01(t *testing.T) {

	r := strings.NewReader(doc)
	crabSubmarines := util.ParseCommaSeparatedInputAsInts(r)
	t.Logf("Crabs %v", crabSubmarines)
	positionWithFuel := make(map[int]map[int]int)

	for _, first := range crabSubmarines {
		positionWithFuel[first] = make(map[int]int)
		for _, second := range crabSubmarines[1:] {
			fuel := util.Abs(first - second)
			positionWithFuel[first][second] = fuel
		}
	}

	t.Logf("Positions %v", positionWithFuel)
}
