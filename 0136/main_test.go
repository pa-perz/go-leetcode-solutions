package main

import "testing"

func Test_singleNumber(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := 1
		nums := []int{2, 2, 1}
		got := singleNumber(nums)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		want := 4
		nums := []int{4, 1, 2, 1, 2}
		got := singleNumber(nums)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("3", func(t *testing.T) {
		want := 1
		nums := []int{1}
		got := singleNumber(nums)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
