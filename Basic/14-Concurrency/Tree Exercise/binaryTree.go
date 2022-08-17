package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	//newChan := make(chan int)
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
	//close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr1 := []int{}
	arr2 := []int{}
	flag := true
	completed := false
	go func() {

		for n := range ch1 {
			arr1 = append(arr1, n)
		}
		//fmt.Println("hello")
		for m := range ch2 {

			arr2 = append(arr2, m)
		}
		newArr1 := (*[10]int)(arr1)
		newArr2 := (*[10]int)(arr2)
		if *newArr1 != *newArr2 {
			//fmt.Println("not equal!")
			//fmt.Println(newArr1)
			//fmt.Println(newArr2)
			flag = false
			completed = true
			return
		}
		flag = true
		completed = true

	}()
	Walk(t1, ch1)
	close(ch1)
	Walk(t2, ch2)
	close(ch2)

	for {
		if completed == true {
			return flag
		}
	}
}

func main() {
	/*c1 := make(chan int)
	go func(){
		for n:= range(c1){
			fmt.Println(n)
		}
	}()
	Walk(tree.New(1), c1)
	fmt.Println(tree.New(1))*/
	tree1 := tree.New(1)
	//tree2:= tree1
	tree2 := tree.New(1)
	same := Same(tree1, tree2)
	fmt.Println(same)
	//fmt.Println(tree1)
	//fmt.Println(tree2)

}
