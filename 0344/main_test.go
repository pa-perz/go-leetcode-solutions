package main

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := []byte{1, 2, 3, 4, 5}
		got := reverseString([]byte{5, 4, 3, 2, 1})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+q want %+q", got, want)
		}
	})
	t.Run("1", func(t *testing.T) {
		want := []byte{1, 2, 3, 4, 5, 6}
		got := reverseString([]byte{6, 5, 4, 3, 2, 1})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+q want %+q", got, want)
		}
	})
}
