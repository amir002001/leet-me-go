package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		midVal := nums[mid]

		if midVal == target {
			return mid
		} else if midVal >= nums[lo] { // left sorted part
			if midVal > target && target >= nums[lo] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if midVal < nums[lo] { // right sorted part
			if midVal < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return -1
}
