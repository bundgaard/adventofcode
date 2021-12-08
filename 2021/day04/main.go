package main

import (
	"2021/util"
	"bufio"
	"log"
	"strings"
)

func main() {
	fd := util.OpenFile("day04.txt")
	br := bufio.NewReader(fd)
	line, err := br.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	drawNumbers := util.ParseCommaSeparatedInputAsInts(strings.NewReader(line))
	boards, err := readBoards(br, 5, 100)
	if err != nil {
		log.Fatal(err)
	}
	var marker [100][5][5]bool
	winningBoards := make(map[int]int)
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
						log.Printf("Board %d won with with full row. Winning #%d", boardIdx+1, n)
						log.Printf("Sum %d", calculateSumUnmarked(boards[boardIdx], marker[boardIdx], n))
						winningBoards[boardIdx]++
					} else if checkColumn(marker[boardIdx]) {
						log.Printf("Board %d won with with full column. Winning #%d", boardIdx+1, n)
						log.Printf("Sum %d", calculateSumUnmarked(boards[boardIdx], marker[boardIdx], n))
						winningBoards[boardIdx]++
					}
					if len(winningBoards) == 100 {
						break BingoLoop
					}
				}
			}
		}
	}

}

func calculateSumUnmarked(board Board, marked [5][5]bool, winningNumber int) int {
	sum := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !marked[y][x] {
				sum += board[y*5+x]
			}
		}
	}
	return sum * winningNumber
}
func checkRow(arr [5][5]bool) bool {
	for y := 0; y < 5; y++ {
		var row int
		for x := 0; x < 5; x++ {
			if arr[y][x] {
				row++
			}
			if row == 5 {
				return true
			}
		}
	}
	return false
}
func checkColumn(arr [5][5]bool) bool {

	for x := 0; x < 5; x++ {
		var marked int
		for y := 0; y < 5; y++ {
			if arr[y][x] {
				marked++
			}
			if marked == 5 {
				return true
			}
		}
	}
	return false
}
