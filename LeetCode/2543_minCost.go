package Alth

//0 1 2
//红 蓝 绿

//每个房间颜色错开
func minCost1(costs [][]int) int {
	dp := costs[0]

	for _, cost := range costs[1:] {
		dpNew := make([]int, 3)
		for j, c := range cost {
			dpNew[j] = min(dp[(j+1)%3], dp[(j+2)%3]) + c
		}
		dp = dpNew
	}

	return min(min(dp[0], dp[1]), dp[2])
}


func minCost(costs [][]int) int {
	dp := make([][3]int, len(costs))
	//初始化

	for i := 0; i < 3; i++ {

		dp[0][i] = costs[0][i]
		for j := 1; j < len(costs); j++ {
			dp[j][0] = min(dp[j-1][1], dp[j-1][2]) + costs[j][0]
			dp[j][1] = min(dp[j-1][0], dp[j-1][2]) + costs[j][1]
			dp[j][2] = min(dp[j-1][0], dp[j-1][1]) + costs[j][2]
		}
	}
	return min(min(dp[len(costs)-1][0], dp[len(costs)-1][1]), dp[len(costs)-1][2])
}
