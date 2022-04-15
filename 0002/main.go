package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := obtainNumber(l1)
	n2 := obtainNumber(l2)

	result := add(n1, n2)
	return createListNode(result)
}

func createListNode(s []int) *ListNode {
	ln := ListNode{}
	ln.Val = int(s[0])
	if len(s) > 1 {
		ln.Next = createListNode(s[1:])
	}
	return &ln
}

func add(n1, n2 []int) []int {
	if len(n1) > len(n2) {
		return sumEach(n1, n2)
	} else {
		return sumEach(n2, n1)
	}
}

func sumEach(large, short []int) []int {
	carry := false
	var result []int
	for i, v := range short {
		sum := v + large[i]
		if carry {
			carry = false
			sum++
		}
		if sum > 9 {
			carry = true
			sum = sum - 10
		}
		result = append(result, sum)
	}
	for _, v := range large[len(short):] {
		sum := v
		if carry {
			carry = false
			sum++
		}
		if sum > 9 {
			carry = true
			sum = sum - 10
		}
		result = append(result, sum)
	}
	if carry {
		result = append(result, 1)
	}

	return result
}

func obtainNumber(l *ListNode) []int {
	var listVal []int
	listVal = append(listVal, l.Val)
	for l.Next != nil {
		l = l.Next
		listVal = append(listVal, l.Val)
	}
	return listVal
}
