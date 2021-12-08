package main

import (
	"2021/util"
	"bufio"
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

	br := bufio.NewScanner(strings.NewReader(doc))

	paths := make(map[int]int)
	for br.Scan() {
		line := br.Text()
		segment := parseLineSegment(line)
		t.Logf("Segment %v", segment)
	}

}

type xy struct {
	X int
	Y int
}
type segment struct {
	Begin xy
	End   xy
}

func parseLineSegment(str string) *segment {
	fields := strings.Fields(str)
	begin := util.ParseCommaSeparatedInputAsInts(strings.NewReader(fields[0]))
	end := util.ParseCommaSeparatedInputAsInts(strings.NewReader(fields[2]))
	return &segment{
		Begin: xy{X: begin[0], Y: begin[1]},
		End:   xy{X: end[0], Y: end[1]},
	}
}
