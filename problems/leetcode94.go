package problems

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0, 100)
	inorderTraversalHelper(root, &ret)
	return ret

}

func inorderTraversalStack(root *TreeNode) []int {
	stack := make([]*TreeNode, 0, 100)
	ret := make([]int, 0, 100)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		root = root.Right
	}
	return ret
}

func inorderTraversalHelper(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	inorderTraversalHelper(root.Left, ret)

	*ret = append(*ret, root.Val)
	inorderTraversalHelper(root.Right, ret)
}
