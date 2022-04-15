package main

import (
	"reflect"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := []string{"1", "2", "Fizz"}
		got := fizzBuzz(3)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+q want %+q", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		want := []string{"1", "2", "Fizz", "4", "Buzz"}
		got := fizzBuzz(5)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+q want %+q", got, want)
		}
	})
	t.Run("3", func(t *testing.T) {
		want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
		got := fizzBuzz(15)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+q want %+q", got, want)
		}
	})
}
