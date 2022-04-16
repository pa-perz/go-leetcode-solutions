package main

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	size := len(nums)
	totalRotation := k % size
	newPos := make(map[int]int)
	for i, n := range nums {
		newPos[(i+totalRotation)%size] = n
	}
	for i := range nums {
		nums[i] = newPos[i]
	}
}
