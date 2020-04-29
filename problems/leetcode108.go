package problems

func sortedArrayToBST(nums []int) *TreeNode {
	return toBstHelper(nums, 0, len(nums)-1)
}

func toBstHelper(nums []int, l, r int) *TreeNode {

	if l > r {
		return nil
	}
	p := (l + r) / 2
	node := &TreeNode{
		Val: nums[p],
	}
	node.Left = toBstHelper(nums, l, p-1)
	node.Right = toBstHelper(nums, p+1, r)
	return node
}
