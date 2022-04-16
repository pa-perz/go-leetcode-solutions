package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	nextLn := head
	for nextLn != nil {
		temp := newHead
		newHead = nextLn
		nextLn = nextLn.Next
		newHead.Next = temp
	}
	return newHead
}
