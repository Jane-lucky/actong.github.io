package main

import "fmt"

func main() {
	ch1 := make(chan string)

	go senData(ch1)

	// //1.循环接受数据
	// for {
	// 	data := <-ch1
	// 	if data == "" {
	// 		break
	// 	}
	// 	fmt.Println("1. 从通道读取的数据为：", data)
	// }

	// //2
	// for {
	// 	data, ok := <-ch1
	// 	fmt.Println(ok)
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("2. 从通道读取的数据为：", data)
	// }
	//3
	for value := range ch1 {
		fmt.Println("3. 从通道读取的数据为：", value)
	}
}

func senData(ch1 chan string) {
	defer close(ch1)
	for i := 0; i < 3; i++ {
		ch1 <- fmt.Sprintf("发送的数据%d\n", i)
	}
	fmt.Println("数据发送完成")
}
