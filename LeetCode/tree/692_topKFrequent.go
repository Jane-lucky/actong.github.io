package tree

import (
	"container/heap"
	"sort"
)

//给定一个单词列表 words 和一个整数 k ，返回前 k 个出现次数最多的单词
//排序+哈希表
func topKFrequent(words []string, k int) []string {
	mmap := make(map[string]int)

	//获取出现的次数
	for _, v := range words {
		mmap[v]++
	}

	restmp := []string{}

	for w := range mmap {
		restmp = append(restmp, w)
	}
	sort.Slice(restmp, func(i, j int) bool {
		s, t := restmp[i], restmp[j]
		return mmap[s] > mmap[t] || mmap[s] == mmap[t] && s < t
	})
	return restmp[:k] //没有按照字典的排序方式来

}

//采用小顶堆的方式
type Item struct {
	word  string
	count int
}
type ItemHeap []Item

func (h ItemHeap) Len() int { return len(h) }
func (h ItemHeap) Less(i, j int) bool {
	if h[i].count != h[j].count {
		return h[i].count < h[j].count
	} else {
		return h[i].word > h[j].word
	}
}
func (h ItemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *ItemHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}
func topKFrequent1(words []string, k int) []string {
	countMap := make(map[string]int)
	for _, word := range words {
		countMap[word]++
	}
	h := &ItemHeap{}
	for w, v := range countMap {
		heap.Push(h, Item{
			word:  w,
			count: v,
		})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]string, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		item := heap.Pop(h).(Item)
		res[i] = item.word
	}
	return res
}
