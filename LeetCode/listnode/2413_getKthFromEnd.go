package listnode

//顺序遍历
func getKthFromEnd(head *ListNode, k int) *ListNode {
	//如果链表的长度小于k，返回一个长度为0的链表

	//先获取链表的长度
	var phead *ListNode

	pre := head
	llen := 0

	for pre != nil {
		llen++
		pre = pre.Next
	}
	if llen < k {
		return phead
	}
	phead = head
	for i := 0; i < llen-k; i++ {
		phead = phead.Next
	}
	return phead
}

//双指针
func getKthFromEnd1(head *ListNode, k int) *ListNode {
	fast := head
	slow := head

	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
