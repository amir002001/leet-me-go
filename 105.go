package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	rootIdx := indexOf(inorder, root.Val)
	root.Left = buildTree(preorder[1:rootIdx+1], inorder[:rootIdx])
	root.Right = buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:])
	return root
}

func indexOf(slice []int, num int) int {
	for k, v := range slice {
		if v == num {
			return k
		}
	}
	return -1
}
