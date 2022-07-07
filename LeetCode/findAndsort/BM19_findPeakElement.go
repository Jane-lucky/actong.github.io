package findandsort

//给定一个长度为n的数组nums，请你找到峰值并返回其索引。数组可能包含多个峰值，在这种情况下，返回任何一个所在位置即可
func findPeakElement(nums []int) int {
	// write code here
	llen := len(nums)

	if llen == 1 {
		return 0
	}

	for i := 1; i < llen-1; i++ {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			return i
		}
	}
	if nums[0] < nums[llen-1] {
		return llen - 1
	} else {
		return 0
	}

}


