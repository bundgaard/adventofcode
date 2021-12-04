package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart01(t *testing.T) {
	doc := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	br := bufio.NewScanner(strings.NewReader(doc))
	for br.Scan() {
		line := br.Text()
		t.Log(line)
	}
}
