package listnode

//递归
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	//每次遍历k次反转，采用count的计数
	curr := head
	count := 0

	for curr != nil && count != k {
		curr = curr.Next
		count++
	}
	if count == k {
		curr = reverseKGroup(curr, k)
		for count != 0 {
			count--
			ptmp := head.Next
			head.Next = curr
			curr = head
			head = ptmp
		}
		head = curr
	}
	return head

}

//模拟
func reverseKGroup1(head *ListNode, k int) *ListNode {
	phead := &ListNode{Next: head}

	pre := phead
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return phead.Next
			}
			nxt := tail.Next
			head, tail := myReverse(head, tail)
			pre.Next = head
			tail.Next = nxt
			pre = tail
			head = tail.Next
		}
	}
	return phead.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	p := head
	pre := tail.Next

	for pre != tail {
		nex := p.Next
		p.Next = pre
		pre = p
		p = nex
	}

	return tail, head
}
