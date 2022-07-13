package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func main() {
// 	input := bufio.NewScanner(os.Stdin)

// 	for input.Scan() { //每调用一次，读取一行
// 		str1 := input.Text()
// 		input.Scan()
// 		s := input.Text()
// 		fmt.Println(StrCount(str1, s))
// 	}
// }

func main() {
	inputreader := bufio.NewReader(os.Stdin)
	str, _ := inputreader.ReadString('\n')
	s, _ := inputreader.ReadString('\n')
	fmt.Println(StrCount(str, s))
}

func StrCount(str, s string) int {
	count := 0
	//将str全部转换成小写
	stmp := []rune(strings.ToLower(str))
	byte := []rune(strings.ToLower(s))
	b := byte[0]
	for _, v := range stmp {
		if v == b {
			count++
		}
	}
	return count
}
