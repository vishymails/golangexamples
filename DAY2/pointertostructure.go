package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books /* Declare Book1 of type Book */
	var Book2 Books /* Declare Book2 of type Book */

	/* book 1 specification */
	Book1.title = "RFRS Programming"
	Book1.author = "Dr. Vishwanath Rao"
	Book1.subject = "Eclipse Plugin Developement"
	Book1.book_id = 6495407

	/* book 2 specification */
	Book2.title = "OO226"
	Book2.author = "Dr. Vishwanath Rao & Jennifer Ball"
	Book2.subject = "Object Oriented Programming"
	Book2.book_id = 6495700

	/* print Book1 info */
	printBook(&Book1)

	/* print Book2 info */
	printBook(&Book2)
}
func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
