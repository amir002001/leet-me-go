package main

func longestConsecutive(nums []int) int {
	presMap := make(map[int]bool)
	maxCount := 0
	for _, num := range nums {
		presMap[num] = true
	}
	for _, num := range nums {
		if _, ok := presMap[num-1]; !ok {
			count := 0
			for {
				count++
				if _, ok := presMap[num+1]; !ok {
					break
				}
				num++
			}
			if count > maxCount {
				maxCount = count
			}
		}
	}
	return maxCount
}
