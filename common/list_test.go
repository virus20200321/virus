package common

import "testing"

func Test_Create(t *testing.T) {
	head := CreateList([]int{1, 2, 3, 4, 5, 6})
	PrintList(head)
}
