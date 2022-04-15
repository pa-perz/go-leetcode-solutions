package main

func searchInsert(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	} else if target <= nums[0] {
		return 0
	} else if target == nums[len(nums)-1] {
		return len(nums) - 1
	} else {
		init, end := 0, len(nums)
		diff := end - init
		var middle int
		for diff > 1 {
			middle = init + (diff)/2
			if target == nums[middle] {
				return middle
			} else if target > nums[middle] {
				if target < nums[middle+1] {
					return middle + 1
				} else {
					init = middle
				}
			} else {
				if target > nums[middle-1] {
					return middle
				} else {
					end = middle
				}
			}
			diff = end - init
		}
		return middle
	}
}
