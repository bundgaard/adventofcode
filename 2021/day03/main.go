package main

import (
	"2021/util"
	"log"
	"strconv"
	"strings"
)

func run(arr []string) map[int][]util.Value {
	bits := make(map[int][]util.Value)

	for _, line := range arr {
		for position, bit := range line {
			bits[position] = append(bits[position], util.Value{Bit: int(bit) - '0'})
		}
	}
	return bits
}

func main() {

	fd := util.OpenFile("day03.txt")
	defer fd.Close()
	lines := util.ParseInput(fd)
	bits := run(lines)
	epsilon, gamma := calculatePowerEfficiencyη(bits)
	log.Println("Part 01")
	log.Printf("%d * %d = %d\n", epsilon, gamma, epsilon*gamma)
	log.Println(strings.Repeat("=", 80))

	log.Println("Part 02")

	oxygenLines := make([]string, len(lines))
	co2Lines := make([]string, len(lines))
	copy(oxygenLines, lines)
	copy(co2Lines, lines)
	oxygenBin := oxygenGeneratorRating(oxygenLines, 0)
	co2Bin := CO2ScrubberRating(co2Lines, 0)
	log.Printf("oxygen %v, co2 %v", oxygenBin, co2Bin)
	oxygen, err := strconv.ParseInt(oxygenBin[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	co2, err := strconv.ParseInt(co2Bin[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Result %d * %d = %d\n", oxygen, co2, oxygen*co2)
}

/////// PART 2
