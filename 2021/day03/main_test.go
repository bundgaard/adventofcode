package main

import (
	"2021/util"
	"sort"
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

func TestPart02(t *testing.T) {
	lines := util.ParseInput(strings.NewReader(doc))
	cp := make([]string, len(lines))
	copy(cp, lines)
	oxygenBin := oxygenGeneratorRating(lines, 0)
	co2Bin := CO2ScrubberRating(cp, 0)
	oxygen, err := strconv.ParseInt(oxygenBin[0], 2, 64)
	if err != nil {
		t.Error(err)
	}
	co2, err := strconv.ParseInt(co2Bin[0], 2, 64)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Result %v * %v = %d", oxygen, co2, oxygen*co2)
	// filter on the most common bit in the first position

}
func calculatePositions(arr []string) map[int]map[byte]int {
	counter := make(map[int]map[byte]int)
	for x := 0; x < len(arr[0]); x++ {
		coll := make(map[byte]int)
		for y := 0; y < len(arr); y++ {
			coll[arr[y][x]]++

		}
		counter[x] = coll
	}
	return counter
}

func recalculateCommons(in map[int]map[byte]int) (leastBits []byte, mostBits []byte) {
	sortedArray := make([]int, 0)
	for k := range in {
		sortedArray = append(sortedArray, k)
	}
	sort.Ints(sortedArray)
	mostBits = make([]byte, 0)
	leastBits = make([]byte, 0)

	for _, k := range sortedArray {
		mostBits = append(mostBits, util.MaxMap(in[k])-'0')
		leastBits = append(leastBits, util.MinMap(in[k])-'0')
	}

	return
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
				t.Logf("breaking at pos %d", pos)
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

func baz(oxylines []string) []string {
	return nil
}

var co2 int64
var oxygen int64

func foo(arr []string, pos int, co2Criteria bool) {
	ones := 0
	for _, elm := range arr {
		if elm[pos] == '1' {
			ones++
		}
	}
	if len(arr) == 1 {
		if co2Criteria {
			co2, _ = strconv.ParseInt(arr[0], 2, 64)
		} else {
			oxygen, _ = strconv.ParseInt(arr[0], 2, 64)
		}
		return
	} else {
		output := make([]string, 0)
		if ones >= len(arr)/2 {
			for _, elm := range arr {
				if co2Criteria {
					if elm[pos] == '1' {
						output = append(output, elm)
					}
				} else {
					if elm[pos] == '0' {
						output = append(output, elm)
					}
				}

			}
		} else {
			for _, elm := range arr {
				if co2Criteria {
					if elm[pos] == '0' {
						output = append(output, elm)
					}
				} else {
					if elm[pos] == '1' {
						output = append(output, elm)
					}
				}

			}
		}

		foo(output, pos+1, co2Criteria)
	}

}

func TestManual(t *testing.T) {
	oxygen, _ := strconv.ParseInt("101001001011", 2, 64)
	co2, _ := strconv.ParseInt("010110110101", 2, 64)

	t.Logf("%d %d = %d", oxygen, co2, oxygen*co2)

}
