package main

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for i := range nums {
		if isDominant(nums, i) {
			return i
		}
	}

	return -1
}

func isDominant(nums []int, index int) bool {
	size := len(nums)
	n := make([]int, size)
	copy(n, nums)
	value := n[index]
	n[index] = n[size-1]
	n = n[:size-1]
	for _, v := range n {
		if v*2 > value {
			return false
		}
	}
	return true
}
