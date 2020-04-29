package problems

/**TreeNode
两种方法都可
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	oldRight := root.Right

	root.Right = root.Left
	root.Left = oldRight

	invertTree(root.Right)
	invertTree(root.Left)
	return root

}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Right)
	right := invertTree(root.Left)
	root.Left = right
	root.Right = left

	return root

}
