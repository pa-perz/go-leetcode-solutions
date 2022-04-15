package main

import "testing"

func Test_reverse(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := 321
		got := reverse(123)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		want := -321
		got := reverse(-123)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("3", func(t *testing.T) {
		want := 21
		got := reverse(120)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("4", func(t *testing.T) {
		want := -1
		got := reverse(-10)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("5", func(t *testing.T) {
		want := 0
		got := reverse(1534236469)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
