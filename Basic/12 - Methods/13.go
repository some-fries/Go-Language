package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	//uncommenting the next line gives an error due to calling a method on a nil interface
	//i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
