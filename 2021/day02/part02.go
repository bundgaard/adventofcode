package main

type SubmarinWithAim struct {
	Depth              int
	HorizontalPosition int
	Aim                int
}

func (sma *SubmarinWithAim) Destination() int {
	return sma.Depth * sma.HorizontalPosition
}

func (sma *SubmarinWithAim) MoveSubmarine(command string) {
	direction, X := parseIntoDirectionAndNumber(command)

	switch direction {
	case Forward:
		sma.HorizontalPosition += X
		sma.Depth += sma.Aim * X
	case Down:
		sma.Aim += X
	case Up:
		sma.Aim -= X
	}
}
