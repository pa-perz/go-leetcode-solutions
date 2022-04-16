package main

func singleNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		found := false
		for j := i + 1; j < len(nums) && !found; j++ {
			if nums[i] == nums[j] {
				found = true
				nums = append(nums[:j], nums[j+1:]...)
				nums = append(nums[:i], nums[i+1:]...)
				i = i - 1
			}
		}
		if !found {
			return nums[i]
		}
	}
	return 0
}
