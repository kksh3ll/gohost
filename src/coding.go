package main

import (
	"fmt"
)

func main() {
	index := 0

	var str string
	fmt.Scanf("%s", &str)
	char := []rune(str)

	switch len(char) {
	case 4:
		index += 1*(int(char[3])-'a') + 1
		index += (25+1)*(int(char[2])-'a') + 1
		index += (25*25+25+1)*(int(char[1])-'a') + 1
		index += (25*25*25 + 25*25 + 25 + 1) * (int(char[0]) - 'a')
	case 3:
		index += (25+1)*(int(char[2])-'a') + 1
		index += (25*25+25+1)*(int(char[1])-'a') + 1
		index += (25*25*25 + 25*25 + 25 + 1) * (int(char[0]) - 'a')
	case 2:
		index += (25*25+25+1)*(int(char[1])-'a') + 1
		index += (25*25*25 + 25*25 + 25 + 1) * (int(char[0]) - 'a')
	case 1:
		index += (25*25*25 + 25*25 + 25 + 1) * (int(char[0]) - 'a')
	default:
		break
	}
	fmt.Printf("%d\n", index)
}
