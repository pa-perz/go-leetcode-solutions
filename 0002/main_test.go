package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 8,
				},
			},
		}
		l1 := &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
				},
			},
		}
		l2 := &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 4,
				},
			},
		}

		got := addTwoNumbers(l1, l2)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("values are not the same")
		}
	})
	t.Run("2", func(t *testing.T) {
		want := &ListNode{
			Val: 0,
		}
		l1 := &ListNode{
			Val: 0,
		}
		l2 := &ListNode{
			Val: 0,
		}

		got := addTwoNumbers(l1, l2)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("values are not the same")
		}
	})
}
