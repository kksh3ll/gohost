package main

import "fmt"

func main() {
	//fmt.Println(int('b') - 'a')

	//reader := bufio.NewReader(os.Stdin)
	//str, _ := reader.ReadString('\n')
	//
	//for _, char := range []rune(str) {
	//	fmt.Println(char)
	//}

	//fmt.Println(len([]rune(str)) - 1)

	var str string
	fmt.Scanf("%s", &str)
	println(str)


	fmt.Println([]rune(str))

}
