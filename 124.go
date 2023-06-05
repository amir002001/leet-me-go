package main

import "math"

func maxPathSum(root *TreeNode) int {
	max := math.MinInt
	helper124(root, &max)
	return max
}

func helper124(root *TreeNode, globalMax *int) int {
	if root == nil {
		return 0
	}
	left := helper124(root.Left, globalMax)
	if left < 0 {
		left = 0
	}
	right := helper124(root.Right, globalMax)
	if right < 0 {
		right = 0
	}
	stage1 := root.Val + left + right
	if stage1 > *globalMax {
		*globalMax = stage1
	}
	var maxKid int
	if right > left {
		maxKid = right
	} else {
		maxKid = left
	}
	return root.Val + maxKid
}
