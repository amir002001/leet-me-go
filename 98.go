package main

import "container/list"

func isValidBST(root *TreeNode) bool {
	inorder := list.New()
	inOrderBb(root, inorder)
	first := true
	var prev int
	for e := inorder.Back(); e != nil; e = e.Prev() {
		if !first && prev >= e.Value.(int) {
			return false
		}
		first = false
		prev = e.Value.(int)
	}
	return true
}

func inOrderBb(root *TreeNode, in *list.List) {
	if root == nil {
		return
	}
	inOrderBb(root.Left, in)
	in.PushFront(root.Val)
	inOrderBb(root.Right, in)
}
