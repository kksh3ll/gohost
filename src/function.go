package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main()  {
	a, b := swap("ligang", "xiwenting")

	fmt.Println(a, b)
}
