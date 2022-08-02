package main

import "fmt"

func main() {
	//Using defined length
	var arr1 = [5]int{1, 2, 3, 4, 5}
	arr2 := [4]int{9, 8, 7, 6}
	fmt.Println("Declaring using defined Length")
	fmt.Println(arr1)
	fmt.Println(arr2)

	//Using inferred length
	var arr3 = [...]string{"This", "is", "GO"}
	arr4 := [...]bool{true, true, false, true}
	fmt.Println("Declaring using inferred length")
	fmt.Println(arr3)
	fmt.Println(arr4)

	//Accessing elements
	fmt.Println("I am accessing the 2nd element of arr3 which is", arr3[2])

	//Change element
	arr2[2] = 99
	fmt.Println("I have changed 3rd element of arr 2 which is now", arr2[2])

	//Assign specific elements during declaration
	arr5 := [5]int{2: 66, 4: 100}
	fmt.Println("arr5 := [5]int{2:66, 4:100} results in", arr5)

	//Length of array
	fmt.Println("Length of arr1 is", len(arr1))

}
