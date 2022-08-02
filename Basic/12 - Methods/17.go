package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (P Person) String() string {
	return fmt.Sprintf("The name is %v and age is (%v)years", P.Name, P.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
