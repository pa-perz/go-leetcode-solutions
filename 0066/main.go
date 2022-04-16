package main

func plusOne(digits []int) []int {
	carry := true
	for i := len(digits) - 1; i >= 0 && carry; i-- {
		value := digits[i] + 1
		if value > 9 {
			value = value % 10
		} else {
			carry = false
		}
		digits[i] = value
	}
	if carry {
		digits = append([]int{1}, digits...)
	}
	return digits
}
