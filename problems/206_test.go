package problems

import "testing"

func TestCreateList(t *testing.T) {
	head := CreateList([]int{1, 2, 3, 4, 5, 6})
	head = reverseBetween(head, 2, 4)
	PrintList(head)
}

func TestReverse(t *testing.T) {
	head := CreateList([]int{1, 2, 3, 4, 5, 6})
	head = reverseList2(head)
	PrintList(head)
}
