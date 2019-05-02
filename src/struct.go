package main

import "fmt"

type books struct {
	title string
	prince float32
	book_id int32
}

func main()  {
	fmt.Println(books{"GO language", 17.32, 913322123})
	fmt.Println(books{"GO language😀", 17.32, 913322123})
	fmt.Println(books{title:"JAVA study", prince:19.43459602, book_id:3210945})
	fmt.Println(books{title:"JAVA study", prince:19.43459602})
	
	//===========================================================================
	var book books
	book.title = "java logic"
	book.prince=30.11111
	book.book_id=231029399


	fmt.Println(book.book_id)
	fmt.Println(book.prince)
	fmt.Println(book.title)

}