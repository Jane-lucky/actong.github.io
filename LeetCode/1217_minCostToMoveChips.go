package Alth

func minCostToMoveChips(position []int) int {
	//判断奇数多还是偶数多
	count1 := 0
	count2 := 0

	for _, v := range position {
		if v&1 == 0 {
			count1++
		} else {
			count2++
		}
	}
	
	return min(count1, count2)
}
