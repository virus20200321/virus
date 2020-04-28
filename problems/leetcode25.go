package problems

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	var a, b = head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverseBetween2Nodes(a, b)
	head.Next = reverseKGroup(b, k)
	return newHead

}

func reverseBetween2Nodes(head1, head2 *ListNode) *ListNode {

	var prev, cur *ListNode = nil, head1
	for cur != head2 {

		n := cur.Next
		cur.Next = prev
		prev = cur
		cur = n
	}
	return prev
}
