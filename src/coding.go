package main

import (
	"fmt"
)

/**
#include <stdio.h>

#define N1 1
#define N2 25
#define N3 (25 * 25)
#define N4 (25 * 25 * 25)

#define C1 N1
#define C2 (N1 + N2)
#define C3 (N1 + N2 + N3)
#define C4 (N1 + N2 + N3 + N4)

int main()
{
char code[5] = {0};
scanf("%s", code);
int index = 0;
switch(strlen(code)){
case 4: index += C1 * (code[3] - 'a') + 1;
case 3: index += C2 * (code[2] - 'a') + 1;
case 2: index += C3 * (code[1] - 'a') + 1;
case 1: index += C4 * (code[0] - 'a');
default: break;
}
printf("%d\n", index);
return 0;
} */
func main()  {
	index:=0
	var code string
	fmt.Scanf("%s", code)
	char:=[]byte(code)
	switch len(char) {
	case 4:index += 1*(char[3]-'a') +1
	case 3:index +=

		
	}
}