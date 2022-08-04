package main

import "fmt"

func pascalTriangle(n int) {
	//counter := 0

	space := " "
	previous := []int{1, 1}
	for i := 1; i < n+1; i++ {
		arr := make([]int, i)
		// if i == 1 {
		// 	fmt.Println(1)
		// } else if i == 2 {
		// 	fmt.Printf("%v %v\n", 1, 1)
		// }
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				arr[j] = 1
			} else {
				arr[j] = previous[j-1] + previous[j]
			}
		}
		//reverseCounter := n
		for k := 0; k < n-i; k++ {
			fmt.Print(space)
		}
		for k := 0; k < i; k++ {
			fmt.Print(arr[k], space)
		}
		//fmt.Println(arr)
		fmt.Println()
		previous = arr
	}
	//counter = +1
}

func main() {
	pascalTriangle(7)
}
