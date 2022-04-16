package main

import "testing"

func Test_reverseVowels(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := "holle"
		got := reverseVowels("hello")
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		want := "leotcede"
		got := reverseVowels("leetcode")
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("3", func(t *testing.T) {
		want := ""
		got := reverseVowels("")
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("4", func(t *testing.T) {
		want := "leoetcede"
		got := reverseVowels("leeetcode")
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("4", func(t *testing.T) {
		want := "Aa"
		got := reverseVowels("aA")
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
