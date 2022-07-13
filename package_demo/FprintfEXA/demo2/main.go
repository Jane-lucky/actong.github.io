package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputreader := bufio.NewReader(os.Stdin)

	input, err := inputreader.ReadString('\n')
	if err == nil {
		fmt.Println(input)
	}
}
