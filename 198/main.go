package main

func main() {
	rob([]int{})
}

func rob(nums []int) int {
	rob1 := 0
	rob2 := 0
	for _, num := range nums {
		tmp := max(num+rob1, rob2)
		rob1 = rob2
		rob2 = tmp
	}
	return rob2
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
