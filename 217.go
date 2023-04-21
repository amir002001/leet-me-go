package main

func containsDuplicate(nums []int) bool {
	dupMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		dupMap[nums[i]] = true
	}
	return len(dupMap) != len(nums)
}
