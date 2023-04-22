package main

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}
	buckets := make([][]int, len(nums)+1)

	for key, value := range freqMap {
		buckets[value] = append(buckets[value], key)
	}
	out := []int{}
	for i := len(buckets) - 1; i >= 0; i-- {
		bucket := buckets[i]
		for _, num := range bucket {
			if k != 0 {
				out = append(out, num)
				k--
			}
		}
	}
	return out
}
