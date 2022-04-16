package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		want := &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		}
		ln := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}
		got := reverseList(ln)
		if !reflect.DeepEqual(got, want) {
			t.Error("values are not equal")
		}
	})
}
