package Alth

type point struct {
	x int
	y int
}

//单遍历
func maxDistance1(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// res = math.Max(float64(res),findNearLands(i,j))
		}
	}
	return res - 1
}

// func findNearLands(i, j int) float64{
// 	// var queue []*point
// }

//多源BFS
func maxDistance(grid [][]int) int {
	//初始化
	var queue []*point

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				queue = append(queue, &point{i, j})
			}
		}
	}

	//如果全是陆地，或者海洋，返回-1
	if len(queue) == 0 || len(queue) == len(grid)*len(grid[0]) {
		return -1
	}

	ans := 0
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	//遍历队列
	for len(queue) > 0 {
		ans++
		tmpQueue := queue
		queue = nil
		for len(tmpQueue) > 0 {
			p := tmpQueue[0]
			tmpQueue = tmpQueue[1:]
			for i := 0; i < len(dir); i++ {
				x := p.x + dir[i][0]
				y := p.y + dir[i][1]
				//判断是否存在
				if x < 0 || x > len(grid) || y < 0 || y > len(grid) || grid[x][y] != 0 {
					continue
				}
				queue = append(queue, &point{x, y})
				grid[x][y] = 2

			}
		}

	}
	return ans - 1
}
