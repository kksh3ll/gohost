package main

import "fmt"

//s := int[]{1, 3, 4}

func main() {
	var number []int

	printSlice(number)

	if (number == nil) {
		fmt.Println("切片是空的")
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(x), cap(x), x)
}
