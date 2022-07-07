package listnode

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var phead *ListNode

	for head != nil {
		ptmp := head.Next
		head.Next = phead
		phead = head
		head = ptmp
	}
	return phead
}


