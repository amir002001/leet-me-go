package main

import "math"

func isBadVersion(version int) bool { return false }

func firstBadVersion(n int) int {
	sqrt := int(math.Sqrt(float64(n)))

	for i := 0; i < sqrt; i++ {
		whatToCheck := (i + 1) * sqrt
		if whatToCheck > n {
			whatToCheck = n
		}
		if isBadVersion(whatToCheck) {
			idx := i*sqrt + 1
			for !isBadVersion(idx) {
				idx++
			}
			return idx
		}
	}

	return -1
}
