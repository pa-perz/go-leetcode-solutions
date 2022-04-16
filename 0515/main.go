package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	tn := []*TreeNode{root}

	for len(tn) != 0 {
		max := math.MinInt32
		level := len(tn)

		for i := 0; i < level; i++ {
			left := tn[i].Left
			if left != nil {
				tn = append(tn, left)
			}
			right := tn[i].Right
			if right != nil {
				tn = append(tn, right)
			}
			if tn[i].Val > max {
				max = tn[i].Val
			}
		}
		result = append(result, max)
		tn = tn[level:]
	}

	return result
}
