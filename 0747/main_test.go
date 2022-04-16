package main

import "testing"

func Test_dominantIndex(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := 1
		got := dominantIndex([]int{3, 6, 1, 0})
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		want := -1
		got := dominantIndex([]int{1, 2, 3, 4})
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("3", func(t *testing.T) {
		want := 0
		got := dominantIndex([]int{1})
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("4", func(t *testing.T) {
		want := 3
		got := dominantIndex([]int{0, 0, 0, 1})
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
