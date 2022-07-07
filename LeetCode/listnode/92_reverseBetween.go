package listnode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here

	//可预设一个空节点，最后返回该节点的next
	rhead := &ListNode{Val: -1}
	rhead.Next = head

	pre := rhead

	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	rightNode := pre
	for i := 0; i < n-m+1; i++ {
		rightNode = rightNode.Next
	}
	leftNode := pre.Next
	cur := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil
	reverse(leftNode)

	pre.Next = rightNode
	leftNode.Next = cur

	return rhead.Next

}

func reverse(head *ListNode) {
	var phead *ListNode

	for head != nil {
		ptmp := head.Next
		head.Next = phead
		phead = head
		head = ptmp
	}
}
