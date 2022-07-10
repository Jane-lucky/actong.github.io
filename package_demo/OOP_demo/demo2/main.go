package main

import "fmt"

type Dog struct {
	name  string
	color string
	age   int
	kind  string
}

func main() {
	//1.实现深拷贝
	d1 := Dog{"豆豆", "黄色", 4, "二哈"}
	fmt.Printf("d1: %T , %v , %p \n", d1, d1, &d1)
	d2 := d1
	fmt.Printf("d2: %T , %v , %p \n", d2, d2, &d2)
	d2.age = 10
	fmt.Printf("d1: %T , %v , %p \n", d1, d1, &d1)
	fmt.Printf("d2: %T , %v , %p \n", d2, d2, &d2)

	//浅拷贝
	d3 := &d1
	d3.name = "球球"
	fmt.Printf("d1: %T , %v , %p \n", d1, d1, &d1)
	fmt.Printf("d3: %T , %v , %p \n", d3, d3, &d3)

}
