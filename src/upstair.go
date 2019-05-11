package main

import "fmt"

func main() {
	sum := climbstairs(9)
	fmt.Println(sum)
}

func climbstairs(n int) int {
	var prev = 0
	var cur = 1
	for i := 1; i <= n; i++ {
		tmp := cur
		cur += prev
		prev = tmp
	}
	return cur
}
