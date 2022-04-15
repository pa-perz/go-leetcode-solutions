package main

import (
	"math"
	"strconv"
)

var minRange = int(math.Pow(-2, 31))
var maxRange = int(math.Pow(2, 31) - 1)

func reverse(x int) int {
	val := []rune(strconv.Itoa(x))
	var start int
	if val[0] == '-' {
		start = 1
	} else {
		start = 0
	}
	nums := val[start:]
	size := len(nums)
	for i, j := 0, size-1; i < (size)/2; i, j = i+1, j-1 {
		buffer := nums[i]
		nums[i] = nums[j]
		nums[j] = buffer
	}

	var sign string
	if start == 1 {
		sign = "-"
	} else {
		sign = ""
	}

	result, _ := strconv.Atoi(sign + string(nums))
	if result > maxRange || result < minRange {
		return 0
	}
	return result
}
