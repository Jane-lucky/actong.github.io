package simulation

func solve(n int, m int, a []int) []int {
	// write code here
	llen := n - m%n
	for i := 0; i < llen; i++ {
		a = append(a, a[i])
	}
	return a[llen:]
}
