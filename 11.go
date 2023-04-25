package main

import "math"

func maxArea(height []int) int {
	max := 0

	i, j := 0, len(height)-1

	for i < j {
		water := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
		if water > max {
			max = water
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return max
}
