package simulation

func rotateMatrix(mat [][]int, n int) [][]int {
	// write code here

	res := [][]int{}

	for i := 0; i < n; i++ {
		ans := []int{}
		for j := n - 1; j >= 0; j-- {
			ans = append(ans, mat[j][i])
		}
		res = append(res, ans)

	}
	return res
}
