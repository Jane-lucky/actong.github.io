package listnode

type ListNode struct {
	Next *ListNode
	Val  int
}

func getLength(head *ListNode) (llen int) {
	for ; head != nil; head = head.Next {
		llen++
	}
	return
}
