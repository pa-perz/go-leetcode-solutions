package main

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	want := []int{0, 1, 9, 16, 100}
	got := sortedSquares([]int{-4, -1, 0, 3, 10})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}
