package Alth

//求最长的斐波那契序列
func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	mmp := make(map[int]int, n)
	for i, x := range arr {
		mmp[x] = i
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i, x := range arr {
		for j := n - 1; j >= 0 && arr[j]*2 > x; j-- {
			if k, ok := mmp[x-arr[j]]; ok {
				dp[j][i] = max(dp[k][j]+1, 3)
				ans = max(ans, dp[j][i])
			}
		}
	}
	return ans
}
