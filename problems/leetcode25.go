package problems

//https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/di-gui-si-wei-ru-he-tiao-chu-xi-jie-by-labuladong/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	b := head
	for i := 0; i < k; i++ {

		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverseAB(head, b)
	//FIXME b is not newhead [a,b) b maybe null
	head.Next = reverseKGroup(b, k)
	return newHead
}

func reverseAB(a, b *ListNode) *ListNode {
	var prev, cur *ListNode = nil, a

	for cur != b {
		n := cur.Next
		cur.Next = prev
		prev = cur
		cur = n

	}
	return prev
}
