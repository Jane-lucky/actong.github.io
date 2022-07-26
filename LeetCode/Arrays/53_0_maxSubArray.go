package arrays

import (
	"fmt"
	"math"
)


//正反两遍扫描，某个点之后s(1,q)<0,maxSubArray6不符合要求
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	llen := len(nums)
	p := -1
	for i := 0; i < llen; i++ {
		if nums[i] > 0 {
			p = i
			break
		}
	}
	if p == -1 {
		max := math.MinInt16
		for i := 0; i < llen; i++ {
			if nums[i] > max {
				max = nums[i]
			}
		}
		return max
	}
	return 0

}

//正反两遍扫描，前提：S(1,q)>0——
func maxSubArray6(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxsum := 0
	llen := len(nums)
	p := -1
	for i := 0; i < llen; i++ {
		if nums[i] > 0 {
			p = i
			break
		}
	}
	if p == -1 {
		max := math.MinInt16
		for i := 0; i < llen; i++ {
			if nums[i] > max {
				max = nums[i]
			}
		}
		return max
	}

	sum := nums[p]
	maxf := sum
	r := 0
	for i := p + 1; i < llen; i++ {
		sum += nums[i]
		if sum > maxf {
			maxf = sum
			r = i
		}
	}
	sum = nums[llen-1]
	maxf = sum
	l := llen - 1
	for i := llen - 2; i >= p; i-- {
		sum += nums[i]
		if sum > maxf {
			maxf = sum
			l = i
		}
	}

	for i := l; i <= r; i++ {
		maxsum += nums[i]
	}
	return maxsum
}

//分治
func maxSubArray5(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return maxSubArraySum(nums, 0, len(nums)-1)
}

func maxSubArraySum(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}
	mid := l + (r-l)/2
	lsum := maxSubArraySum(nums, l, mid)
	rsum := maxSubArraySum(nums, mid+1, r)
	asum := maxCrossingSum(nums, l, mid, r)
	return max(lsum, max(rsum, asum))
	// return asum
}
func maxCrossingSum(nums []int, l, mid, r int) int {
	sum := 0
	leftnum := math.MinInt16

	for i := mid; i >= l; i-- {
		sum += nums[i]
		if sum > leftnum {
			leftnum = sum
		}
	}
	sum = 0
	rightsum := math.MinInt16
	for i := mid + 1; i <= r; i++ {
		sum += nums[i]
		if sum > rightsum {
			rightsum = sum
		}
	}
	return leftnum + rightsum
}

func maxSubArray4(nums []int) int {
	return get(nums, 0, len(nums)-1).msum
}

func pushUp(l, r Status) Status {
	iSum := l.isum + r.isum
	lSum := max(l.lsum, l.isum+r.lsum)
	rSum := max(r.rsum, r.isum+l.rsum)
	mSum := max(max(l.msum, r.msum), l.rsum+r.lsum)
	return Status{lSum, rSum, mSum, iSum}
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lsub := get(nums, l, m)
	rsub := get(nums, m+1, r)
	return pushUp(lsub, rsub)
}

type Status struct {
	lsum, rsum, msum, isum int
}

//动态规划
func maxSubArray3(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//二重循环
func maxSubArray2(nums []int) int {
	sum := math.MinInt16
	for i := 0; i < len(nums); i++ {
		stmp := 0
		for j := i; j < len(nums); j++ {
			stmp += nums[j]
			if stmp > sum {
				sum = stmp
			}
		}

	}
	return sum
}

//三重循环
func maxSubArray1(nums []int) int {
	sum := math.MinInt16
	left, right := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			stmp := 0
			for k := i; k <= j; k++ {
				stmp += nums[k]
			}
			if stmp > sum {
				sum = stmp
				left = i
				right = j
			}
		}
	}
	fmt.Println(nums[left : right+1])
	return sum
}
