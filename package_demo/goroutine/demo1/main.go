package main

import (
	"fmt"
	"time"
)

// func Hello() {
// 	fmt.Println("Hello function")
// }
// func main() {
// 	go Hello()
// 	time.Sleep(50 * time.Microsecond)
// 	fmt.Println("main function")
// }

func main() {
	go running()
	var input string
	fmt.Scanln(&input)
}

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}
