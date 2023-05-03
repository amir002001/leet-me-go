package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if isSameTree(root, subRoot) {
		return true
	}
	if root == nil {
		return false
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
