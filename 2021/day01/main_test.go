package main

import "testing"

func TestSlidingSum(t *testing.T) {
	measurements := []int{
		199, 200, 208, 210, 200, 207, 240, 269, 260, 263,
	}

	slidingSums := countSlidingWindow(measurements, 3)
	t.Logf("Sums %v", slidingSums)
	if len(slidingSums) != 8 {
		t.Fail()
	}

	count := countIncrement(slidingSums)
	t.Logf("Counts %v", count)

}
