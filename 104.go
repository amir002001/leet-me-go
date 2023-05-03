package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	var max int
	if leftMax > rightMax {
		max = leftMax
	} else {
		max = rightMax
	}
	return max + 1
}
