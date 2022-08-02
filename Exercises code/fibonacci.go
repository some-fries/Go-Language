package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	previousOfPrevious := 0
	previous := 0
	counter := 0

	return func() int {
		if previous == 0 && counter == 0 {
			counter += 1
			return 0
		} else if previous == 0 && counter == 1 {
			previous += 1
			counter += 1
			return 1
		}

		//fib := previous + previousOfPrevious
		temp := previousOfPrevious
		previousOfPrevious = previous
		previous = previous + temp
		//fmt.Println("Previous is", previous)
		//currentFib := fib
		return previous //currentFib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
