package main

import (
	"2021/util"
	"bufio"
	"log"
	"strings"
)

var points = map[int32]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var pairs = map[int32]int32{
	'(': ')',
	'{': '}',
	'<': '>',
	'[': ']',
}

var doc = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`

func part01() {
	fd := util.OpenFile("day10.txt")
	br := bufio.NewScanner(fd)
	corruptions := make([]int32, 0)
	for br.Scan() {
		line := br.Text()
		corruption, ok := findCorruptedNavigation(line)
		if ok {
			corruptions = append(corruptions, corruption)
		}
	}
	total := 0
	for _, x := range corruptions {
		total += points[x]
	}
	log.Printf("Navigational corruptions %d %q total=%d", len(corruptions), corruptions, total)
}
func main() {

	br := bufio.NewScanner(strings.NewReader("[({(<(())[]>[[{[]{<()<>>"))

	incompletes := make([]string, 0)
	for br.Scan() {
		line := br.Text()
		incomplete, ok := findIncompleteLine(line)
		if !ok {
			continue
		}
		incompletes = append(incompletes, incomplete)
	}

	closers := make([][]int32, 0)
	for _, incomplete := range incompletes {
		missing := countClosing(incomplete)
		closers = append(closers, missing)
	}

	log.Printf("Closers %q", closers)

	total := calculateAutocomplete(closers[0])
	log.Printf("Part02 Total %d", total)

}

func calculateAutocomplete(xs []int32) int {
	points := map[int32]int{')': 1, ']': 2, '}': 3, '>': 4}
	total := 0
	for idx, x := range xs {
		point := points[x]
		log.Printf("#%03d\t %c -> %d", idx, x, point)
		total *= 5
		log.Printf("#%03d\t %c -> %d, total = %d", idx, x, point, total)
		total += points[x]
	}
	return total
}
func countClosing(line string) []int32 {
	level := 0
	queue := make([]int32, 0)
	missing := make([]int32, 0)
	for idx, current := range line {
		switch current {
		case '{':
			fallthrough
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '<':
			queue = append(queue, current)
			level++
		case '}':
			fallthrough
		case ')':
			fallthrough
		case ']':
			fallthrough
		case '>':
			last := queue[util.MinMax(len(queue)-1, 0)]
			log.Printf("#%03d\tlast=%c current=%c", idx, last, current)
			queue = queue[:util.MinMax(len(queue)-1, 0)]
			level--
		}
	}

	for _, open := range queue {
		missing = append(missing, pairs[open])
	}
	log.Printf("not closed %q", queue)
	return missing
}
func findIncompleteLine(line string) (string, bool) {
	stack := make([]int32, 0)
	level := 0

	for idx, current := range line {
		switch current {
		case '{':
			fallthrough
		case '(':
			fallthrough
		case '<':
			fallthrough
		case '[':
			level++
			stack = append(stack, current)
			// fmt.Printf("#%d\t%s %c \n", level, strings.Repeat("\t", level), current)
		case ']':
			fallthrough
		case '}':
			fallthrough
		case '>':
			fallthrough
		case ')':
			// Incomplete?
			last := stack[util.MinMax(len(stack)-1, 0)]
			// can we close the pair
			// log.Printf("last=%q current=%q stack=%q rest=%q", last, current, stack, line[idx:])
			if pairs[last] == current {
				// we can close so pop the value
				stack = stack[0:util.MinMax(len(stack)-1, 0)]
				level--
				//				log.Printf("popped stack %q", stack)
			} else {
				return "", false
			}
			if idx == len(line)-1 && len(stack) > 0 {
				//log.Printf("incomplete %s", line)
				return line, true
			}

		}
	}
	return line, true
}

func findCorruptedNavigation(line string) (int32, bool) {
	pairs := map[int32]int32{
		'(': ')',
		'{': '}',
		'<': '>',
		'[': ']',
	}
	stack := make([]int32, 0)
	level := 0
	for idx, current := range line {
		switch current {
		case '{':
			fallthrough
		case '(':
			fallthrough
		case '<':
			fallthrough
		case '[':
			level++
			stack = append(stack, current)
			// fmt.Printf("#%d\t%s %c \n", level, strings.Repeat("\t", level), current)
		case ']':
			fallthrough
		case '}':
			fallthrough
		case '>':
			fallthrough
		case ')':
			// Incomplete?
			last := stack[util.MinMax(len(stack)-1, 0)]
			// can we close the pair
			// log.Printf("last=%q current=%q stack=%q rest=%q", last, current, stack, line[idx:])
			if pairs[last] == current {
				// we can close so pop the value
				stack = stack[0:util.MinMax(len(stack)-1, 0)]
				level--
				//				log.Printf("popped stack %q", stack)
			} else {
				return current, true
			}
			if idx == len(line)-1 && len(stack) > 0 {
				log.Printf("incomplete %s", line)
				return 0, false
			}

		}
	}
	return 0, false
}
