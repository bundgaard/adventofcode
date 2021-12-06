package main

import (
	"2021/util"
	"strconv"
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

func TestPart02A(t *testing.T) {
	lines := util.ParseInput(util.OpenFile("/home/dbundgaard/code/adventofcode/2021/day03.txt"))
	oxylines := make([]string, len(lines))
	copy(oxylines, lines)
	co2lines := make([]string, len(lines))
	copy(co2lines, lines)
	for {
		if len(oxylines) == 1 {
			break
		}
		for pos := range oxylines[0] {
			counter := make(map[byte]int)
			for _, line := range oxylines {
				counter[line[pos]]++
			}
			var b byte
			if counter['0'] == counter['1'] {
				b = '1'
			} else {
				b = util.MaxMap(counter)
			}
			candidates := make([]string, 0)
			for _, line := range oxylines {
				if line[pos] == b {
					candidates = append(candidates, line)
				}
			}
			oxylines = candidates
		}
	}

CO2Scrubber:
	for {
		if len(co2lines) == 1 {
			break
		}
		for pos := range co2lines[0] {
			counter := make(map[byte]int)
			for _, line := range co2lines {
				counter[line[pos]]++
			}
			var b byte
			if counter['0'] == counter['1'] {
				b = '0'
			} else {
				b = util.MinMap(counter)
			}
			candidates := make([]string, 0)
			if len(co2lines) <= 2 {
				t.Logf("breaking at pos %d; %d", pos, len(co2lines))
				break CO2Scrubber
			}
			for _, line := range co2lines {
				if line[pos] == b {
					candidates = append(candidates, line)
				}
			}
			co2lines = candidates
		}
	}

	t.Logf("OxyLines %v %d", oxylines, len(oxylines))
	oxygen, _ = strconv.ParseInt(oxylines[0], 2, 64)
	t.Logf("co2lines %v %d", co2lines, len(co2lines))
	co2, _ = strconv.ParseInt(co2lines[0], 2, 64)
	t.Logf("co2 %d, oxygen %d %d", co2, oxygen, co2*oxygen)

}

var co2 int64
var oxygen int64
