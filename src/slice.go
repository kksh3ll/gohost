package main

import "fmt"

//s := int[]{1, 3, 4}

func main() {
	var number []int

	//number = []int {1,2,3,4,5}

	printSlice(number)

	if (number == nil) {
		fmt.Println("切片是空的")
	} else {
		fmt.Println("not null")
	}

	numbers1 := make([]int, 0, 7)

	fmt.Println(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(x), cap(x), x)
}
