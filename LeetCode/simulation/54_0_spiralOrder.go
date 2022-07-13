package simulation

func spiralOrder1(matrix [][]int) []int {
	// write code here
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, colums := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, colums)
	}

	var (
		total          = rows * colums
		order          = make([]int, total)
		row, column    = 0, 0
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directionIndex = 0
	)
	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		nextrow, nexcolumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextrow < 0 || nextrow >= rows || nexcolumn < 0 || nexcolumn >= colums || visited[nextrow][nexcolumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}

//模拟打圈
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, colums             = len(matrix), len(matrix[0])
		order                    = make([]int, rows*colums)
		index                    = 0
		left, right, top, bottom = 0, colums - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		//从左到右
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		//从上到下
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			//从右到左
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			//从下到上
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		top++
		right--
		bottom--

	}
	return order
}
