package main

import "fmt"

func main() {

	for i := 0; i < 11; i++ {
		//fmt.Printf("%d ", fibonacci(i))
		fmt.Printf("%d ", fib_dp(i))
	}

}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func fib_dp(n int) int {
	var f = [40]int{}
	f[0] = 0
	f[1] = 1
	if n < 2 {
		return n
	}
	for i := 2; i < n+1; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n-1] +f[n-2]
}
