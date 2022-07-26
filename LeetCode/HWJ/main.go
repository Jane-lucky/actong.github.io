package main

import (
	"fmt"
)

func main() {

	L := make([]int, 10)
	n, h, a, b := 0, 0, 0, 0
	fmt.Scanf("%d %d", &n, &h)
	L[h] = 0
	for n = n - 1; n > 0; n-- {
		fmt.Scanf("%d %d", &a, &b)
		L[a] = L[b]
		L[b] = a
	}
	fmt.Println(L)
	// skip or delete
	fmt.Scanf("%d", &a)

	for h > 0 {
		if h == a {
			h = L[a]
		} else {
			fmt.Printf("%d ", h)
			h = L[h]
		}
	}
}
