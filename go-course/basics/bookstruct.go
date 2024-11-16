package main

import "fmt"

type Book struct {
	Title string
	Author string
	Pages int
}

func bookstruct() {
	harry_potter := Book{
		Title: "Harry Potter and the Chamber of Secrets",
		Author: "J.K. Rowling",
		Pages: 340,
	}
	fmt.Printf("Title: %s, Author: %s, Pages: %d\n", harry_potter.Title, harry_potter.Author, harry_potter.Pages);
}