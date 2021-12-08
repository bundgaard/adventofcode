package main

import (
	"2021/util"
	"bufio"
	"strings"
	"testing"
)

var doc = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
 `

func TestPart01(t *testing.T) {
	br := bufio.NewReader(strings.NewReader(doc))
	line, err := br.ReadString('\n')
	if err != nil {
		t.Error(err)
	}
	drawNumbers := util.ParseCommaSeparatedInputAsInts(strings.NewReader(line))

	boards, err := readBoards(br, 5, 3)
	if err != nil {
		t.Error(err)
	}
	var marker [3][5][5]bool
	var winners [3]bool
	_ = winners
BingoLoop:
	for _, n := range drawNumbers {
		for idx := range boards[0] {
			for boardIdx := range boards {
				x := idx % 5
				y := idx / 5
				m := boards[boardIdx][idx]
				if m == n {
					marker[boardIdx][y][x] = true
					if checkRow(marker[boardIdx]) {
						t.Logf("Board %d won with with full row. Winning #%d", boardIdx+1, n)
						t.Logf("Sum %d", calculateSumUnmarked(boards[boardIdx], marker[boardIdx], n))

					} else if checkColumn(marker[boardIdx]) {
						t.Logf("Board %d won with with full column. Winning #%d", boardIdx+1, n)
						t.Logf("Sum %d", calculateSumUnmarked(boards[boardIdx], marker[boardIdx], n))
						break BingoLoop
					}
				}

			}
		}
	}

}
