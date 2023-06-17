package main

func main() {
	_ = climbStairs(10)
}

func climbStairs(n int) int {
	var helper func(n int) int

	memo := make(map[int]int)

	helper = func(n int) int {
		if n < 0 {
			return 0
		}
		if n == 0 {
			return 1
		}
		if value, ok := memo[n]; ok {
			return value
		}
		val := helper(n-1) + helper(n-2)
		memo[n] = val
		return val
	}

	return helper(n)
}
