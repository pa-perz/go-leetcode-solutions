package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := true
		got := isPalindrome(121)
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		want := false
		got := isPalindrome(-121)
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("3", func(t *testing.T) {
		want := false
		got := isPalindrome(10)
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
