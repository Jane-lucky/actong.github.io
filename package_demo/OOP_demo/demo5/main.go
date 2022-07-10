package main

import (
	"fmt"
	"math"
)

func main() {
	//匿名函数
	res := func(a, b float64) float64 {
		return math.Pow(a, b)
	}(2, 3)
	fmt.Println(res)

	//匿名结构体
	addr := struct {
		province, city string
	}{"陕西省", "西安市"}
	fmt.Println(addr)

}
