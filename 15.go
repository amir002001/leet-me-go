package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	indexMap := make(map[int]int)
	for i, num := range nums {
		indexMap[num] = i
	}

	outMap := make(map[[3]int]bool)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			complement := -(nums[i] + nums[j])
			if iComp, ok := indexMap[complement]; ok && iComp != j && iComp != i {
				ints := [3]int{nums[j], nums[i], complement}
				sort.Ints(ints[:])
				outMap[ints] = true
			}
		}
	}

	out := [][]int{}

	for key := range outMap {
		what := key[:]
		out = append(out, what)
	}
	fmt.Println(outMap)
	return out
}
