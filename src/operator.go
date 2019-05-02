package main

import "fmt"

func main ()  {
	var a int  = 4
	var b int32
	var c float32
	var ptr *int


	/* 运算符实例 */

	fmt.Printf("1. a = %T\n", a)
	fmt.Printf("2. b = %T\n", b)
	fmt.Printf("3. c = %T\n", c)
	fmt.Printf("4. ptr = %T\n", ptr)

	ptr = &a

	fmt.Printf("----------------")
	fmt.Printf("ptr = %d\n", ptr)
	fmt.Printf("ptr's value = %d", *ptr)
}