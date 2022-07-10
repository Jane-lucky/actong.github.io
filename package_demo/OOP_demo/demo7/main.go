package main

import "fmt"

type Phone interface {
	call()
}
type AndroidPhone struct {
}
type Iphone struct {
}

func (a AndroidPhone) call() {
	fmt.Println("安卓手机")
}

func (i Iphone) call() {
	fmt.Println("苹果手机")
}

func main() {
	var phone Phone
	phone = new(AndroidPhone)
	fmt.Printf("%T  ,%v  ,%p   \n", phone, phone, &phone)
	phone.call()
	phone = AndroidPhone{}
	fmt.Printf("%T  ,%v  ,%p   \n", phone, phone, &phone)
	phone.call()
	phone = new(Iphone)
	fmt.Printf("%T  ,%v  ,%p   \n", phone, phone, &phone)
	phone.call()
	fmt.Printf("%T  ,%v  ,%p   \n", phone, phone, &phone)
	phone = Iphone{}
	fmt.Printf("%T  ,%v  ,%p   \n", phone, phone, &phone)
	phone.call()
	fmt.Printf("%T  ,%v  ,%p   \n", phone, phone, &phone)
}
