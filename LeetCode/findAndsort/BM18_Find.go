package findandsort

func Find(target int, array [][]int) bool {
	// write code here
	for i := 0; i < len(array); i++ {
		if len(array[i]) == 0 {
			continue
		}
		if array[i][0] > target {
			continue
		} else if array[i][len(array[i])-1] < target {
			continue
		} else {
			for j := 0; j < len(array[i]); j++ {
				if array[i][j] == target {
					return true
				}
			}
		}
	}
	return false
}
