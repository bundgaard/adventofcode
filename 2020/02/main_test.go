package main

import "testing"

func TestExample(t *testing.T) {

	t.Run("abcde", func(t *testing.T) {
		in := requirement{Letter: rune('a'), Min: 1, Max: 3, Password: "abcde"}
		if !verifyPasswordPosition(in) {
			t.Fail()
		}

	})

	t.Run("cdefg", func(t *testing.T) {
		in2 := requirement{Letter: rune('b'), Min: 1, Max: 3, Password: "cdefg"}
		if verifyPasswordPosition(in2) {
			t.Fail()
		}
	})

	t.Run("ccccccccc", func(t *testing.T) {
		in2 := requirement{Letter: rune('c'), Min: 2, Max: 9, Password: "ccccccccc"}
		if verifyPasswordPosition(in2) {
			t.Fail()
		}
	})

}
