package main

import "fmt"

type Flower struct {
	name  string
	color string
}

func main() {
	//1. 作为参数
	f1 := Flower{"玫瑰", "红"}
	fmt.Printf("f1 : %T , %v , %p \n", f1, f1, &f1)

	//将结构体对象作为参数
	changeInfo(f1)
	fmt.Printf("changeInfo之后的信息 : %T , %v , %p \n", f1, f1, &f1)

	//结构体指针作为参数
	changeInfo2(&f1)
	fmt.Printf("changeInfo2之后的信息 : %T , %v , %p \n", f1, f1, &f1)

	//2 结构体作为返回值
	//结构体对象作为返回值
	f2 := getFlower()
	fmt.Printf("getFlower外f2的信息 : %T , %v , %p \n", f2, f2, &f2)
	f3 := getFlower()

	fmt.Printf("getFlower外f3的信息 : %T , %v , %p \n", f3, f3, &f3)

	//结构体对象指针作为返回值
	f4 := getFlower1()
	fmt.Printf("getFlower外f4的信息 : %T , %v , %p \n", f4, f4, &f4)
	f5 := getFlower1()

	fmt.Printf("getFlower外f5的信息 : %T , %v , %p \n", f5, f5, &f5)
}

func getFlower() (f Flower) {
	f = Flower{"牡丹", "白"}
	fmt.Printf("getFlower中的信息 : %T , %v , %p \n", f, f, &f)
	return
}

func getFlower1() (f *Flower) {
	f = &Flower{"牡丹", "白"}
	fmt.Printf("getFlower1中的信息 : %T , %v , %p \n", f, f, &f)
	return
}
func changeInfo(f Flower) {
	f.name = "月季"
	fmt.Printf("changeInfo中的信息 : %T , %v , %p \n", f, f, &f)
}

func changeInfo2(f *Flower) {
	f.color = "粉红"
	fmt.Printf("changeInfo2中的信息 : %T , %v , %p , %p \n", f, f, f, &f)
}
