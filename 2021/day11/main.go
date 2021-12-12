package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

var (
	doc = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`
)

func main() {
	br := bufio.NewScanner(strings.NewReader(doc))

	consortium := make([][]int, 0)
	for br.Scan() {
		line := br.Text()
		row := make([]int, 0)

		for _, octopus := range line {
			n := int(octopus - '0')
			row = append(row, n)
		}
		consortium = append(consortium, row)
	}

	step(consortium)
	log.Printf("Consortium %v", consortium)

	display(consortium)
}

func display(consortium [][]int) {
	for y := 0; y < len(consortium); y++ {
		for x := 0; x < len(consortium[0]); x++ {
			fmt.Printf("%d", consortium[y][x])
		}
		fmt.Println()
	}
}
func step(consortium [][]int) {
	flash := 0
	for y := 0; y < len(consortium); y++ {
		for x := 0; x < len(consortium[0]); x++ {
			if consortium[y][x]-1 == 0 {
				flash++
				// consortium[y-1][x]--
				// nw  n ne -1,-1  0,-1 1,-1
				// w c e    -1, 0, 0, 0 1, 0
				// sw s se   1, 1  0, 1 1, 1
			}
			consortium[y][x]--
		}
	}

	log.Printf("total flashes %03d", flash)
}
