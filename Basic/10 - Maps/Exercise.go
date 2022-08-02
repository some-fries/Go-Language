package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	countsMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		elem, ok := countsMap[words[i]]

		if ok == false {
			countsMap[words[i]] = 1
		} else {
			fmt.Println("elem =", elem, "ok =", ok, "\n")
			countsMap[words[i]] += 1
		}
	}
	return countsMap //map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
	fmt.Println(WordCount)
}
