package main

func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		leftSum := 0
		rightSum := 0
		for j := i - 1; j >= 0; j-- {
			leftSum += nums[j]
		}
		for j := i + 1; j < len(nums); j++ {
			rightSum += nums[j]
		}
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}
