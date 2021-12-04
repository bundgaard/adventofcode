package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestPilot(t *testing.T) {
	doc := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	br := bufio.NewScanner(strings.NewReader(doc))
	submarine := SubmarinWithAim{}
	for br.Scan() {
		line := br.Text()
		submarine.MoveSubmarine(line)
	}

	t.Logf("Location of submarine is %d %d %d\n", submarine.HorizontalPosition, submarine.Depth, submarine.Destination())

}
