package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	arr := strings.Split(input.Text(), " ")
	// 	fmt.Println(len(arr[len(arr)-1]))
	// }
	inputreader := bufio.NewReader(os.Stdin)
	data,_,_:=inputreader.ReadLine()
	fmt.Println(Length(string(data)))

}

func Length(s string) int {
	arr := strings.Split(s, " ")
	return len(arr[len(arr)-1])
}
