package Alth

import "math/rand"

type Solution struct {
	b2w   map[int]int
	bound int
}

func Constructor(n int, blacklist []int) Solution {
	m := len(blacklist)

	bound := n - m

	black := map[int]bool{}

	for _, b := range blacklist {
		if b >= bound {
			black[b] = true
		}
	}

	b2w := make(map[int]int, m-len(black))
	w := bound
	for _, b := range blacklist {
		if b < bound {
			for black[w] {
				w++
			}
			b2w[b] = w
			w++
		}
	}
	return Solution{b2w, bound}
}

func (this *Solution) Pick() int {
	x := rand.Intn(this.bound)
	if this.b2w[x] > 0 {
		return this.b2w[x]
	}
	return x
}
