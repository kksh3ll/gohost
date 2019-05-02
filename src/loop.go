package main

import "fmt"

func main() {

	//这是main函数的局部变量
	var i = 18

	fmt.Println("--------------start-------------")
	// i变量是for循环中的局部变量
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("--------------end---------------")

	fmt.Println(i)
}
