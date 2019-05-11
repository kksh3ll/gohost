package main

import "fmt"

func main() {
	twowei()
}

func twowei() {
	var array = [][]int{{1, 2, 4}, {2, 3, 4}, {4, 5, 6}}
	//array

	for index, value := range array {
		fmt.Println(index, value)

		for index1, value1 := range value {
			fmt.Println(index1, value1)
		}
	}
}
