package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, nil, nil)
}

func isValidBSTHelper(root, min, max *TreeNode) bool {
	//root logical
	if root == nil {
		return true
	}
	if min != nil && root.Val > min.Val {
		return false
	}
	if max != nil && root.Val < max.Val {
		return false
	}
	//traverse
	return isValidBSTHelper(root.Left, min, root) && isValidBSTHelper(root.Right, root, max)
}

func commonProcess(t *TreeNode) {
	//root logical

	//traverse

}
