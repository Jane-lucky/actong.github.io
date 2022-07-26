package dp

//递归
func putApple(m, n int) int {
	if m < 0 || n < 0 {
		return 0
	}
	if m == 1 || n == 1 || m == 0 {
		return 1
	}
	return putApple(m-n, n) + putApple(m, n-1)
}

//动态规划
func putApple1(m, n int) int {
	dp := make([][]int, m+1)
	//初始化 没有苹果，只有一种摆法
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//如果苹果数小于盘子数
			//忽略一个盘子，在n-1个放盘子，一直递推下去
			if i < j {
				dp[i][j] = dp[i][j-1]
			} else {
				//苹果数大于或者等于盘子数，两种情况
				dp[i][j] = dp[i][j-1] + dp[i-j][j]
			}
		}
	}
	return dp[m][n]
}
