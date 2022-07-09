package listnode

//顺序合并
func mergeKLists(lists []*ListNode) *ListNode {
	var list *ListNode

	//顺序遍历合并
	for i := 0; i < len(lists); i++ {
		list = mergeTwoLists(list, lists[i])
	}
	return list
}

//双指针合并&&分治合并
func mergeKLists1(lists []*ListNode) *ListNode {
	return mergeList(lists, 0, len(lists)-1)
}

func mergeList(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	return mergeTwoLists(mergeList(lists, left, mid), mergeList(lists, mid+1, right))
}

//小顶堆
