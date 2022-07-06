package Alth

import (
	"math"
	"sort"
)

/**
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。

思路：给nums进行排序
-1,2,1,-4
-4 -1 2 1
*/

func ThreeSumClosest1(nums []int, target int) int {
	//排序加双指针
	sort.Ints(nums)

	var (
		llen = len(nums)
		best = math.MaxInt32
	)

	//根据差值来更新答案
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}
	//双指针
	for i := 0; i < llen; i++ {
		//排除和上一个元素相同
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, llen-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			update(sum)
			if sum > target {
				k0 := k - 1
				//移动到不等于k的坐标下
				for j < k0 && nums[k] == nums[k0] {
					k0--
				}
				k = k0
			} else {
				j0 := j + 1
				//右移值与下标j不相等的坐标
				for j0 < k && nums[j] == nums[j0] {
					j0++
				}
				j = j0
			}
		}
	}
	return best

}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
func ThreeSumClosest(nums []int, target int) int {
	llen := len(nums)

	if llen == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	sort.Ints(nums)
	res := make(map[int]int)
	for i := 0; i < llen-2; i++ {
		tmp := nums[i : i+3]
		var sum, r int

		for _, v := range tmp {
			sum += v
		}
		//判断距离
		if target > sum {
			r = target - sum
		} else {
			r = sum - target
		}
		res[r] = sum

	}

	var r []int
	//遍历循环map
	for k := range res {
		r = append(r, k)
	}
	sort.Ints(r)
	return res[r[0]]
}
