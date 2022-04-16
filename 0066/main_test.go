package main

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := []int{1, 2, 4}
		got := plusOne([]int{1, 2, 3})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+q want %+q", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		want := []int{4, 3, 2, 2}
		got := plusOne([]int{4, 3, 2, 1})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+q want %+q", got, want)
		}
	})
	t.Run("3", func(t *testing.T) {
		want := []int{1, 0}
		got := plusOne([]int{9})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+q want %+q", got, want)
		}
	})
	t.Run("4", func(t *testing.T) {
		want := []int{5, 0, 0, 0}
		got := plusOne([]int{4, 9, 9, 9})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+q want %+q", got, want)
		}
	})
}
