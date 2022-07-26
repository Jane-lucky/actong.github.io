package listnode

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func Construct() {
	//链表构建
	input := bufio.NewScanner(os.Stdin)
	for {
		input.Scan()
		if len(input.Text()) < 1 {
			break
		}
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		head := &ListNode{}
		cur := head
		//采用头插法
		for i := 0; i < n; i++ {
			if i == 0 {
				head.Val, _ = strconv.Atoi(strings.Split(input.Text(), " ")[i])
				continue
			}
			p := &ListNode{}
			p.Val, _ = strconv.Atoi(strings.Split(input.Text(), " ")[i])
			cur.Next = p
			cur = p

		}

	}
}
