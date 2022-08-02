package main

import "fmt"

func main() {
	//Create slice
	slice1 := []int{}
	slice2 := []int{1, 2, 3, 4, 5}

	fmt.Println("slice1 =", slice1)
	fmt.Println("slice2 len:", len(slice2))
	fmt.Println("slice2 cap:", cap(slice2))

	//Creat array
	arr1 := [5]int{9, 8, 7, 6, 5}
	//Extract a slice from the array
	slice3 := arr1[1:3]
	fmt.Println("slice3:", slice3, "\n\tlen:", len(slice3), "\n\tcap:", cap(slice3))
	//The cap is different from the original array since the slice starts from 2nd element of the array and can grow to the end of the array

	//Slice creation with make function make([]type, len, cap)
	slice4 := make([]int, 6, 10)
	fmt.Println("slice4:", slice4, "\n\tlen:", len(slice4), "\n\tcap:", cap(slice4))

	//If cap is not specified then it is equal to the len
	slice5 := make([]int, 7)
	fmt.Println("slice5:", slice5, "\n\tlen:", len(slice5), "\n\tcap:", cap(slice5))

}
