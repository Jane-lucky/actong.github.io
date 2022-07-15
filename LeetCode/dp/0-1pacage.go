package dp

//其中N为物品，W为承重，data为物品的重量及对于的价值
func package0And1(N, W int, data [][]int) int {
	dp := make([][]int, N+1)

	//动态初始化
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}

	//循环遍历
	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if data[i-1][0] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-data[i-1][0]]+data[i-1][1])
			}
		}
	}
	return dp[N][W]
}
func package0And2(N, W int, data [][]int) [][]int {
	dp := make([][]int, N+1)

	//动态初始化
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}

	//循环遍历
	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if data[i-1][0] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-data[i-1][0]]+data[i-1][1])
			}
		}
	}
	return dp
}

func trackback(dp, data [][]int, W int) []bool {
	x := make([]bool, len(dp)-1)

	for i := len(dp) - 1; i > 1; i-- {
		if dp[i][W] == dp[i-1][W] {
			x[i-1] = false
		} else {
			x[i-1] = true
			W -= data[i-1][0]
		}
	}
	return x
}
