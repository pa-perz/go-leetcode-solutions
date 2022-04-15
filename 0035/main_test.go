package main

import "testing"

func Test_searchInsert(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := 2
		nums := []int{1, 3, 5, 6}
		target := 5
		got := searchInsert(nums, target)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		want := 1
		nums := []int{1, 3, 5, 6}
		target := 2
		got := searchInsert(nums, target)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("3", func(t *testing.T) {
		want := 4
		nums := []int{1, 3, 5, 6}
		target := 7
		got := searchInsert(nums, target)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("4", func(t *testing.T) {
		want := 0
		nums := []int{1, 3, 5, 6}
		target := 0
		got := searchInsert(nums, target)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("5", func(t *testing.T) {
		want := 0
		nums := []int{1, 3}
		target := 1
		got := searchInsert(nums, target)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("6", func(t *testing.T) {
		want := 1
		nums := []int{1, 3}
		target := 3
		got := searchInsert(nums, target)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("7", func(t *testing.T) {
		want := 2
		nums := []int{1, 3, 5}
		target := 4
		got := searchInsert(nums, target)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
