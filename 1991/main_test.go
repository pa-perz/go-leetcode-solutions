package main

import "testing"

func Test_pivotIndex(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := 3
		got := pivotIndex([]int{1, 7, 3, 6, 5, 6})
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		want := -1
		got := pivotIndex([]int{1, 2, 3})
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("3", func(t *testing.T) {
		want := 0
		got := pivotIndex([]int{2, 1, -1})
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
