package main

func reverseString(s []byte) []byte {
	size := len(s)
	var buffer byte
	for i, j := 0, size-1; i < size/2; i, j = i+1, j-1 {
		buffer = s[i]
		s[i], s[j] = s[j], buffer
	}
	return s
}
