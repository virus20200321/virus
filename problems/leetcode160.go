package problems

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	oldHeadA := headA
	oldHeadB := headB
	for headA != headB {
		if headA == nil {
			headA = oldHeadB
		} else {
			headA = headA.Next
		}
		if headB == nil {
			headB = oldHeadA

		} else {
			headB = headB.Next
		}

	}
	return headA
}
