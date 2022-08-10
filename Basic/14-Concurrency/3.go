package main

import "fmt"

func main() {
	c := make(chan string, 2) // making a buffered channel allowing the code to move forward in main without causing the channel to block
	c <- "hello"
	c <- "brudda"
	// c <- "third" //this will cause the channel to become full and thus block as it exceeds the buffer capacity
	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
