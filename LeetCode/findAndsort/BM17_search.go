package findandsort

//前提条件：元素升序，返回下标
//顺序遍历
func search(nums []int, target int) int {
	// write code here
	llen := len(nums)

	if llen == 0 {
		return -1
	}

	for i := 0; i < llen; i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

//二分查找
func search1(nums []int, target int) int {
	llen := len(nums)

	if llen == 0 {
		return -1
	}
	left, right := 0, llen-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
