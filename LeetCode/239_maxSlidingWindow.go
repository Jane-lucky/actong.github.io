package Alth

import (
	"container/heap"
	"sort"
)

/**
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。

*/

//滑动窗口为k，则遍历整数数组len(nums)-k+1次，分别获取k个字符的最大值，时间复杂度为O(nk),超出时间限制
//问题：当采用sort给截取数组排序时，会影响原始数组的排序状态
//如果采用max函数的方式，会产生超出时间限制错误
//[]int{7, 2, 4}, 2
func MaxSlidingWindow(nums []int, k int) []int {

	var res []int
	var tmp []int
	if nums == nil || len(nums) < k {
		return res
	}
	for i := 0; i < len(nums)-k+1; i++ {
		tmp = nums[i : i+k]
		sort.Ints(tmp)
		res = append(res, tmp[len(tmp)-1])
	}
	return res
}

/**
滑动窗口：相邻的滑动窗口，公用k-1个元素，而只有一个元素时变化的。
优先队列：将向右开始滑动的时候，可以把新元素优先放入队列,堆顶就是滑动窗口的最大值
在优先队列中存储二元组(nums,index)，表示nums在数组中下标的位置index
*/

var a []int

type hp struct {
	sort.IntSlice
}

func (h *hp) Less(i, j int) bool { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func MaxSlidingWindow1(nums []int, k int) []int {
	a=nums
	q:=&hp{make([]int, k)}

	for i:=0;i<k;i++{
		q.IntSlice[i]=i
	}

	heap.Init(q)

	n:=len(nums)

	ans:=make([]int, 1,n-k+1)

	ans[0]=nums[q.IntSlice[0]]

	for i:=k;i<n;i++{
		heap.Push(q,i)
		for q.IntSlice[0] <= i-k{
			heap.Pop(q)
		}

		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}
