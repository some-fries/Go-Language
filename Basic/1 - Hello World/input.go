package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var input string
	input := bufio.NewReader(os.Stdin)
	line, _ := input.ReadString('\n')
	fmt.Println(line)
	// fmt.Scanln(&input)
	// fmt.Println(input)
}
