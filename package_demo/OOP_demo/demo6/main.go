package main

import "fmt"

type Address struct {
	province, city string
}

type Person struct {
	name string
	age  int
	addr *Address
}

//聚合关系
func main() {
	//模拟结构体之间的聚合关系
	p := Person{}
	p.name = "Lily"
	p.age = 15
	//赋值方式1
	addr := Address{"陕西省", "西安市"}
	p.addr = &addr
	fmt.Println(p)
	fmt.Println("姓名:", p.name, "年龄：", p.age, "地址：", p.addr.province, p.addr.city)
}
