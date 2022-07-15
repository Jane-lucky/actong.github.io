package dp

func minDistance(word1 string, word2 string) int {
	lw1 := len(word1)
	lw2 := len(word2)

	dp := make([][]int, lw1+1)

	//初始化dp
	for i := 0; i < lw1+1; i++ {
		dp[i] = make([]int, lw2+1)
		dp[i][0] = i
	}
	for i := 0; i < lw2+1; i++ {
		dp[0][i] = i
	}

	//遍历判断最短距离
	for i := 1; i < lw1+1; i++ {
		for j := 1; j < lw2+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			}
		}
	}
	return dp[lw1][lw2]
}
