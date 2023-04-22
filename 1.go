package main

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for i, num := range nums {
		indexMap[num] = i
	}
	for i, num := range nums {
		complement := target - num
		if j, ok := indexMap[complement]; ok && j != i {
			return []int{j, i}
		}
	}
	return []int{}
}
