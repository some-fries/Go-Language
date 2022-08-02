package main

import (
	"fmt"
)

func FirstReverse(str string) string {

	// code goes here
	newStr := "dsad"
	for i := 0; i < len(str); i++ {
		//newStr = newStr + str[i]
		//fmt.Println(str)
	}
	fmt.Println(string(str[3]))
	str = newStr
	return str

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	//fmt.Println(
	FirstReverse("hello brudda")

}
