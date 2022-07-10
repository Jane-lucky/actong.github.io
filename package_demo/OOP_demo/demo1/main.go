package main

import "fmt"

type Emp struct {
	name string
	age  int
	sex  byte
}

func main() {
	emp1 := new(Emp)
	fmt.Println(emp1.age)
}
