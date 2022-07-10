package main

import (
	"fmt"
	"time"
)

func main() {
	go printNums()
	go printLetter()

	time.Sleep(3 * time.Second)
	fmt.Println("main over >>>>>")
}

func printNums() {
	for i := 0; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Print(i)
	}
}

func printLetter() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}
