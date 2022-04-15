package main

import "testing"

func Test_IsValid(t *testing.T) {
	t.Run("first string", func(t *testing.T) {
		want := true
		got := isValid("()")

		if want != got {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("second string", func(t *testing.T) {
		want := true
		got := isValid("()[]{}")

		if want != got {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("third string", func(t *testing.T) {
		want := false
		got := isValid("(]")

		if want != got {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("fourth string", func(t *testing.T) {
		want := true
		got := isValid("{[]}")

		if want != got {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
