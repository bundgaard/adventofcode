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
				// 5 1 5
				// 5 5 5

				for i := -1; i < 2; i++ {
					for j := -1; j < 2; j++ {

						log.Printf("%d,%d", i, j)

					}
				}
				/*consortium[y-1][x-1]-- // nw
				consortium[y-1][x+0]-- // n
				consortium[y-1][x+1]-- // ne

				consortium[y][x-1]-- // w
				consortium[y][x+0]-- // c
				consortium[y][x+1]-- // e

				consortium[y+1][x-1]-- // sw
				consortium[y+1][x+0]-- // s
				consortium[y+1][x+1]-- // se
				*/
			} else {
				consortium[y][x]--
			}

		}
	}

	log.Printf("total flashes %03d", flash)
}
