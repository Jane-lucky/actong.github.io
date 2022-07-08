package listnode

//删除倒数第n个节点

//顺序遍历
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	llen := getLength(head)

	var res *ListNode
	if llen < n || n < 0 {
		return res
	}

	dum := &ListNode{Next: head}
	curr := dum
	for i := 0; i < llen-n; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return dum.Next
}

//栈的方式
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	//
	nodes := []*ListNode{}
	dum := &ListNode{head, 0}
	for node := dum; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dum.Next

}
