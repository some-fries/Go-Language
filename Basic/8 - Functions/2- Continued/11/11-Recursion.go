package main

import "fmt"

func factorial(x int) int {
	if x == 0 || x == 1 {
		return 1
	} else {
		return factorial(x-1) * x
	}
}

func main() {
	fmt.Println(factorial(5))
}
