package problems

func IsSameTree(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return IsSameTree(root1.Left, root2.Left) && IsSameTree(root1.Right, root2.Right)

}
