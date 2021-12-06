package main

import (
	"2021/util"
	"log"
)

func main() {
	fd := util.OpenFile("day06.txt")
	lanternFishSchool := util.ParseCommaSeparatedInputAsInts(fd)
	lanternFishSchool = simulate(256, lanternFishSchool)
	log.Println("Part 01")
	log.Printf("Total fish %d", len(lanternFishSchool))
}

func simulate(days int, school []int) []int {
	for i := 0; i < days; i++ {
		school = spawnNewLanternFish(school)
	}
	return school
}
func spawnNewLanternFish(school []int) []int {
	newFish := 0
	for idx := range school {
		lanternFish := school[idx]

		if lanternFish == 0 {
			newFish++
		}

		lanternFish = (lanternFish - 1)
		if lanternFish < 0 {
			lanternFish = 6
		}
		school[idx] = lanternFish
	}

	for i := newFish; i > 0; i-- {
		school = append(school, 8)
	}

	return school
}
