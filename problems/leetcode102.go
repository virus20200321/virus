package problems

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for {
		curQueue := make([]*TreeNode, 0)
		curRow := make([]int, 0)
		for len(queue) > 0 {
			top := queue[0]
			curRow = append(curRow, top.Val)
			queue = queue[1:]
			if top.Left != nil {
				curQueue = append(curQueue, top.Left)
			}
			if top.Right != nil {
				curQueue = append(curQueue, top.Right)
			}
		}
		ret = append(ret, curRow)
		queue = curQueue
		if len(queue) == 0 {
			break
		}

	}
	return ret
}
