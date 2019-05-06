package main

import (
	"container/list"
	"fmt"
)

func main() {

	var n int
	var m int
	fmt.Scanf("%d, %d", &n, &m)

	for i := 1; i < n; i++ {
		if i>n-i {
			break
		}
		m = n - i
		fmt.Printf("%d %d\n", i, m)
	}
	fmt.Println(n)

	example()


}

func example()  {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
