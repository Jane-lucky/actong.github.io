package listnode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dum1 := &ListNode{}
	dum := dum1
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			dum.Next = l1
			l1 = l1.Next
		} else {
			dum.Next = l2
			l2 = l2.Next
		}
		dum = dum.Next
	}

	if l1 != nil {
		dum.Next = l1
	}
	if l2 != nil {
		dum.Next = l2
	}
	return dum1.Next
}
