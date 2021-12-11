package main

import (
	"2021/util"
	"bufio"
	"log"
	"strings"
	"testing"
)

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

func TestPart01(t *testing.T) {
	br := bufio.NewScanner(strings.NewReader(`{<[[]]>}<{[{[{[]{()[[[]`))

	corruptions := make([]string, 0)
	for br.Scan() {
		line := br.Text()
		corruptions = append(corruptions, findCorruptedNavigation(line))
	}

	t.Logf("Navigational corruptions %d %q", len(corruptions), corruptions)
}

/*
{
	<
		[
			[]
		]
	>
}
<
	{
		[
			{
				[
					{
						[]
							{
								()
									[
										[
											[]
*/
func findCorruptedNavigation(line string) string {
	pairs := map[int32]int32{
		'(': ')',
		'{': '}',
		'<': '>',
		'[': ']'}
	stack := make([]int32, 0)

	for idx, b := range line {
		switch b {
		case '{':
			fallthrough
		case '(':
			fallthrough
		case '<':
			fallthrough
		case '[':
			stack = append(stack, b)
		case ']':
			fallthrough
		case '}':
			fallthrough
		case '>':
			fallthrough
		case ')':
			log.Printf("stack %q", stack)
			if len(stack) == 0 {
				continue
			}
			last := stack[util.Max(util.Min(len(stack)-1, 0), 0)]
			log.Printf("#%03d ---  last=%q current=%q", idx, last, b)
			log.Printf("stack %q", stack)

			if pairs[last] == b {
				// remove last
				log.Printf("stack=%d stack-1=%d", len(stack), len(stack)-1)
				stack = stack[0:util.Min(0, len(stack)-1)]
			} else {
				return line
			}

		}
	}
	return ""
}
