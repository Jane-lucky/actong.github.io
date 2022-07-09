package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

//小顶堆和大顶堆
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	//pop和push采用指针接受，是因为他们需要修改切片的长度，不仅仅是他们的长度
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//样例
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	heap.Push(h, 3)
	fmt.Printf("minim:%d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
