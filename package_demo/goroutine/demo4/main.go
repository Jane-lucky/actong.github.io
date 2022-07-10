package main

import "fmt"

func main() {
	var ch1 chan int

	ch1 = make(chan int)
	fmt.Printf("%T\n", ch1)

	ch2 := make(chan bool)

	//匿名Goroutine
	go func() {
		data, ok := <-ch1
		if ok {
			fmt.Println("子routine取到的数值为：", data)
		}
		ch2 <- true
	}()
	ch1 <- 10
	//忽略数据接收，目的是为了阻塞Goroutine
	//防止主函数Goroutine退出导致匿名函数的Goroutine退出
	<-ch2
	fmt.Println("main over")
}
