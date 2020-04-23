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

func reverseList2(head *ListNode) *ListNode {
	var top *ListNode = nil
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = top
		top = cur
		cur = next
	}

	return top
}

var after *ListNode = nil

func reverseN(head *ListNode, n int) *ListNode {
	if n == 0 {
		after = head.Next
		return head
	}

	cur := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = after
	return cur

}

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	if m == 1 {
		return reverseN(head, n-m)
	}

	newNext := reverseBetween(head.Next, m-1, n-1)
	head.Next = newNext
	return head
}

func reverseList0422(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := reverseList0422(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}

var theEnd *ListNode = nil

//1,2,3,4  1,3,2,4
/* This function reverse N times*/
func reverseHelper0422(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}

	newHead := reverseHelper0422(head.Next, n-1)
	head.Next.Next = head
	head.Next = theEnd
	return newHead
}
func reverseBetween0422(head *ListNode, m, n int) *ListNode {
	if m == 1 { // e.g. n=4, m = 2 ; m=1,n=3
		newHead := reverseHelper0422(head, n-1)
		return newHead
	}

	cur := reverseBetween0422(head.Next, m-1, n-1)
	head.Next = cur
	return head
}
