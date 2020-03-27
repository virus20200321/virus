package review

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	cur := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur

}

func reverseList2(head *ListNode) *ListNode {
	prev := head
	if prev == nil || prev.Next == nil {
		return prev
	}
	cur := prev.Next
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev.Next = nil
		prev = cur
		cur = next
	}
	return prev
}
