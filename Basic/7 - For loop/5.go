package main

import "fmt"

func main() {
	cars := [5]string{"ferrari", "lamborgini", "bentley", "porche", "mercedes"}

	//first value is index and second is value
	//If only one value given then it is index
	for _, name := range cars {
		fmt.Println(name)
	}
}
