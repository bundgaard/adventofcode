package main

import (
	"2021/util"
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func createSegments(r io.Reader) []segment {

	segments := make([]segment, 0)

	br := bufio.NewScanner(r)
	for br.Scan() {
		line := br.Text()
		segment := parseLineSegment(line)
		// t.Logf("Segment %v", segment)
		// right or left

		segments = append(segments, *segment)
	}
	return segments
}

func findSizeOfMap(segments []segment) (xMax int, yMax int) {

	for _, segment := range segments {
		xMax = util.Max(xMax, util.Max(segment.Begin.X, segment.End.X))
		yMax = util.Max(yMax, util.Max(segment.Begin.Y, segment.End.Y))

	}
	return
}

func createDangerousMap(xMax, yMax int) [][]int {

	sketch := make([][]int, yMax+1)
	for y := 0; y < yMax+1; y++ {
		sketch[y] = make([]int, xMax+1)
		for x := 0; x < xMax+1; x++ {
			sketch[y][x] = 0
		}
	}
	return sketch
}
func main() {

	fd := util.OpenFile("day05.txt")
	defer func() {
		if err := fd.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	segments := createSegments(fd)
	xMax, yMax := findSizeOfMap(segments)
	log.Printf("found max size for x and y %d %d", xMax, yMax)
	sketch := createDangerousMap(xMax, yMax)

	lookForDangerousAreas(segments, sketch, true)

	// displayDangerousAreas(sketch)
	log.Printf("Number of dangerous area %3d", calculateDangerousAreas(sketch))
}
func lookForDangerousAreas(segments []segment, sketch [][]int, diagonal bool) {

	for _, segment := range segments {
		if !diagonal && segment.Begin.X != segment.End.X && segment.Begin.Y != segment.End.Y {
			continue
		}
		dx := segment.End.X - segment.Begin.X
		dy := segment.End.Y - segment.Begin.Y

		if dy == 0 && dx > 0 {
			horizontal(sketch, segment.Begin.X, segment.End.X, segment.Begin.Y)
		}
		if dy == 0 && dx < 0 {
			horizontal(sketch, segment.End.X, segment.Begin.X, segment.Begin.Y)
		}
		if dx == 0 && dy > 0 {
			vertical(sketch, segment.Begin.Y, segment.End.Y, segment.Begin.X)
		}
		if dx == 0 && dy < 0 {
			vertical(sketch, segment.End.Y, segment.Begin.Y, segment.Begin.X)
		}
		if dx != 0 && dy != 0 {
			px := segment.Begin.X
			py := segment.Begin.Y
			for {
				sketch[py][px]++

				if px == segment.End.X && py == segment.End.Y {
					break
				}
				if dx <= 0 {
					px--
				} else {
					px++
				}
				if dy <= 0 {
					py--
				} else {
					py++
				}
			}
		}
	}
}

func calculateDangerousAreas(sketch [][]int) (sum int) {

	for y := 0; y < len(sketch); y++ {
		for x := 0; x < len(sketch[y]); x++ {
			if sketch[y][x] >= 2 {
				sum++
			}
		}
	}
	return
}
func displayDangerousAreas(sketch [][]int) {
	for y := 0; y < len(sketch); y++ {
		for x := 0; x < len(sketch[y]); x++ {
			point := sketch[y][x]
			if point == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", point)
			}
		}
		fmt.Println()
	}
}
func horizontal(sketch [][]int, x1, x2, y int) {
	for x := x1; x <= x2; x++ {
		sketch[y][x]++
	}
}

func vertical(sketch [][]int, y1, y2, x int) {
	for y := y1; y <= y2; y++ {
		sketch[y][x]++
	}
}

type xy struct {
	X int
	Y int
}

func (e xy) String() string {
	return fmt.Sprintf("XY(%d,%d)", e.X, e.Y)
}

type segment struct {
	Begin xy
	End   xy
}

func parseLineSegment(str string) *segment {
	fields := strings.Fields(str)
	begin := util.ParseCommaSeparatedInputAsInts(strings.NewReader(fields[0]))
	end := util.ParseCommaSeparatedInputAsInts(strings.NewReader(fields[2]))
	return &segment{
		Begin: xy{X: begin[0], Y: begin[1]},
		End:   xy{X: end[0], Y: end[1]},
	}
}
