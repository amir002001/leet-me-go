package main

import "math"

func maxProfit(prices []int) int {
	max := math.MinInt32
	i, j := 0, 1

	for j < len(prices) {
		profit := prices[j] - prices[i]
		if profit > max {
			max = profit
		}
		if prices[j] < prices[i] {
			i = j
		}
		j++
	}
	if max < 0 {
		return 0
	}
	return max
}
