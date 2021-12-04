package main

type SubmarineWithoutAim struct {
	HorizontalPosition int
	Depth              int
}

func (sm *SubmarineWithoutAim) MoveSubmarine(command string) {
	dir, depth := parseIntoDirectionAndNumber(command)
	switch dir {
	case Forward:
		sm.HorizontalPosition += depth
	case Up:
		sm.Depth -= depth
	case Down:
		sm.Depth += depth
	}
}

func (sm *SubmarineWithoutAim) Destination() int {
	return sm.HorizontalPosition * sm.Depth
}
