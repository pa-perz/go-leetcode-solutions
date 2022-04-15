package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

const bad = 4

func isBadVersion(version int) bool {
	return version <= bad
}

func firstBadVersion(n int) int {
	init, end := 0, n
	diff := end - init
	for diff > 1 {
		middle := init + (end-init)/2
		if isBadVersion(middle) {
			end = middle
		} else {
			init = middle
		}
		diff = end - init
	}
	return end
}
