package main

import (
	"2021/util"
	"log"
	"strings"
)

func main() {
	fd := util.OpenFile("day06.txt")
	lanternFish := util.ParseCommaSeparatedInputAsInts(fd)
	lanternFishSchool := loadSchool(lanternFish)
	simulate(80, &lanternFishSchool)
	log.Println("Part 01")
	log.Printf("Total fish %d", sumSchool(&lanternFishSchool))
	log.Println(strings.Repeat("=", 80))
	lanternFishSchool = loadSchool(lanternFish)
	simulate(256, &lanternFishSchool)
	log.Printf("lantern fish after 256 days %d", sumSchool(&lanternFishSchool))

}
func sumSchool(school *School) int {
	result := 0
	for _, val := range school {
		result += val
	}
	return result
}
func loadSchool(school []int) (pool School) {
	for _, fish := range school {
		pool[fish]++
	}
	log.Printf("loadSchool %v", pool)
	return
}
func simulate(days int, school *School) {
	for i := 0; i < days; i++ {
		spawnNewLanternFish(school)
	}
}

type School [9]int

func spawnNewLanternFish(school *School) {
	breeding := school[0]
	for i, fish := range school[1:] {
		school[i] = fish
	}
	school[8] = breeding
	school[6] += breeding
}
