package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, tmp int
	fmt.Scan(&N)
	res := []int{}
	for i := 0; i < N; i++ {
		fmt.Scanln(&tmp)
		res = append(res, tmp)
	}
	sort.Ints(res)
	for _, v := range res {
		fmt.Println(v)
	}
}
