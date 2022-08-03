package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l List[any]) Print() {
	newL := l
	for {
		if newL.next != nil {
			fmt.Printf("%v -> ", newL.val)
			newL = *newL.next
		} else if newL.next == nil {
			fmt.Println(newL.val)
			return
		}
	}
}

func (l *List[any]) Insert(v any) {
	newL := l
	for {
		if newL.next == nil {
			newL.next = &List[any]{nil, v}
			return
			//newL.val = v
		} else {
			newL = newL.next
		}
	}
}

func main() {
	l1 := List[int]{&List[int]{&List[int]{nil, 13}, 7}, 4}
	//fmt.Println(l1)
	l1.Print()
	l1.Insert(20)
	l1.Insert(99)
	l1.Insert(200)
	l1.Insert(2)
	l1.Insert(9)
	l1.Insert(33)
	l1.Print()

}
