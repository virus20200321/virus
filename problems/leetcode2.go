func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := make([]*ListNode, 0), make([]*ListNode, 0)
	for l1 != nil {
		stack1 = append(stack1, l1)
		l1 = l1.Next

	}

	for l2 != nil {
		stack2 = append(stack2, l2)
		l2 = l2.Next
	}
	dump := &ListNode{}
	head := dump
	add := 0
	for len(stack1) > 0 && len(stack2) > 0 {
		top1 := stack1[0]
		stack1 = stack1[1:]
		top2 := stack2[0]
		stack2 = stack2[1:]
		sum := top1.Val + top2.Val + add
		add = sum / 10
		sum = sum % 10
		//      fmt.Println("sum = ", sum )

		head.Next = &ListNode{Next: nil, Val: sum}
		head = head.Next
	}
	var rest []*ListNode
	if len(stack1) > 0 {
		rest = stack1
	} else {
		rest = stack2
	}
	//  fmt.Println("add", add)
	for len(rest) > 0 {
		top := rest[0]
		rest = rest[1:]
		sum := top.Val + add
		// fmt.Println("sum = ", top.Val,"+",add )
		add = sum / 10

		sum = sum % 10

		head.Next = &ListNode{Next: nil, Val: sum}
		head = head.Next
	}
	// fmt.Println("add = ", add)
	if add > 0 {
		head.Next = &ListNode{Next: nil, Val: add}
		head = head.Next
	}
	return dump.Next

}