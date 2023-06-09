package main

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	curr := root
	for len(stack) != 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return curr.Val
		}
		curr = curr.Right
	}
	return -1
}
