package main

import "fmt"

func main() {

	for i := 0; i < 11; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}

}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
