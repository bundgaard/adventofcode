package main

import (
	"strings"
	"testing"
)

var doc = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`

func TestPart01(t *testing.T) {

	r := strings.NewReader(doc)
	segments := createSegments(r)
	xMax, yMax := findSizeOfMap(segments)
	sketch := createDangerousMap(xMax, yMax)
	lookForDangerousAreas(segments, sketch, true)
	sum := calculateDangerousAreas(sketch)
	expected := 12
	displayDangerousAreas(sketch)
	if sum != expected {
		t.Errorf("expected %d got %d", expected, sum)
	}
}
