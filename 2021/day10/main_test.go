package main

import (
	"2021/util"
	"bufio"
	"testing"
)

func TestPart01(t *testing.T) {
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

	if len(corruptions) != 5 {
		t.Errorf("expected 5 corruptions. got %d", len(corruptions))
	}
	total := 0
	for _, x := range corruptions {
		total += points[x]
	}
	t.Logf("Navigational corruptions %d %q total=%d", len(corruptions), corruptions, total)
}
