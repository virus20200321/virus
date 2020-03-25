package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(src []int) *ListNode {
	dummy := &ListNode{}
	ptr := dummy
	for _, v := range src {
		ptr.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		ptr = ptr.Next
	}
	return dummy.Next
}

func PrintList(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	PrintList(head.Next)
}
