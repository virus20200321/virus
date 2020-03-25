package problems

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}

/*
Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
*/

var after *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		after = head.Next
		return head
	}
	cur := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = nil
	return cur

}
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//到达反转位
	if m == 1 {
		return reverseN(head, n)
	}

	cur := reverseBetween(head.Next, m-1, n-1)
	head.Next = cur
	return head

}
