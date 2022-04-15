package main

import "testing"

func Test_search(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := 4
		nums := []int{-1, 0, 3, 5, 9, 12}
		target := 9
		got := search(nums, target)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		want := -1
		nums := []int{-1, 0, 3, 5, 9, 12}
		target := 2
		got := search(nums, target)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
