package problems

func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0, 100)
	preorderTraversalHelper(root, &ret)
	return ret
}

func preorderTraversalHelper(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	*ret = append(*ret, root.Val)
	preorderTraversalHelper(root.Left, ret)
	preorderTraversalHelper(root.Right, ret)

}

func preorderTraversalStack(root *TreeNode) []int {
	ret := make([]int, 0, 100)
	stack := make([]*TreeNode, 0, 100)
	if root == nil {
		return ret
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		out := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if out.Right != nil {
			stack = append(stack, out.Right)
		}
		if out.Left != nil {
			stack = append(stack, out.Left)
		}
		ret = append(ret, out.Val)
	}

	return ret
}
