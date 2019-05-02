package main

import "fmt"

func main()  {
	var a = 10

	fmt.Printf("变量的地址是：%x\n", &a)

	var ptr *int

	ptr = &a

	fmt.Println(ptr)

	fmt.Println(*ptr)

	//nil pointer
	var p *string

	fmt.Println(p)

}