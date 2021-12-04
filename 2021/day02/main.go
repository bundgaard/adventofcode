package main

import (
	"2021/util"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Submarine interface {
	MoveSubmarine(string)
	Destination() int
}

type Direction string

const (
	Forward Direction = "forward"
	Down    Direction = "down"
	Up      Direction = "up"
)

func parseIntoDirectionAndNumber(line string) (Direction, int) {
	words := strings.Split(line, " ")
	dir := Direction(words[0])
	n, err := strconv.Atoi(words[1])
	if err != nil {
		log.Fatal(err)
	}
	return dir, n

}

func main() {

	fd := util.OpenFile("day02.txt")
	defer fd.Close()

	br := bufio.NewScanner(fd)
	submarineWithAim := SubmarinWithAim{}
	submarineWithoutAim := SubmarineWithoutAim{}
	for br.Scan() {
		line := br.Text()
		submarineWithoutAim.MoveSubmarine(line)
		submarineWithAim.MoveSubmarine(line)
	}

	fmt.Println("Part 01")
	fmt.Printf("Location of submarine is %d %d %d\n", submarineWithoutAim.HorizontalPosition, submarineWithoutAim.Depth, submarineWithoutAim.Destination())
	fmt.Println("Part 02")
	fmt.Printf("Location of submarine is %d %d %d\n", submarineWithAim.HorizontalPosition, submarineWithAim.Depth, submarineWithAim.Destination())
}
