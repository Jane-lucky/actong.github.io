package Alth

/**
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

{name: "1", args: args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, want: 49},
		{name: "2", args: args{[]int{1, 1}}, want: 1},
		{name: "3", args: args{[]int{}}, want: 0},
*/
func MaxArea(height []int) int {
	llen := len(height)

	res := 0
	j, k := 0, llen-1
	for j < k {
		var sum int
		if height[j] < height[k] {
			sum = (k - j) * height[j]
			j++

		} else {
			sum = (k - j) * height[k]
			k--

		}

		if sum > res {
			res = sum
		}

	}
	return res
}
