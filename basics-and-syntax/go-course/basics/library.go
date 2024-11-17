package main

import (
	"fmt"
)

type Library struct {
	Books map[int]Book
}

func librarygoprogram() {
	library := Library{
		Books: map[int]Book{
			1: {
				Title: "book1title",
				Author: "book1author",
				Pages: 222,
			},
			2: {
				Title: "book2title",
				Author: "book2author",
				Pages: 333,
			},
		},
	}

	for id, book := range library.Books {
		fmt.Printf("ID: %d\nBook title: %s\tBook author: %s\tBook pages: %d\n", id, book.Title, book.Author, book.Pages);
	}

	fmt.Println();
	book, err := getBookById(1, library);
	if err != nil {
		fmt.Println("Error:", err);
	} else {
		fmt.Printf("Book title: %s\nBook author: %s\nBook pages: %d\n", book.Title, book.Author, book.Pages);
	}

	fmt.Println();
	book, err = getBookById(3, library);
	if err != nil {
		fmt.Println("Error:", err);
	} else {
		fmt.Printf("Book title: %s\nBook author: %s\nBook pages: %d\n", book.Title, book.Author, book.Pages);
	}

}

func getBookById(id int, library Library) (Book, error) {
	book, exists := library.Books[id];
	if exists {
		return book, nil;
	} else {
		return Book{}, fmt.Errorf("Book with id '%d' not found", id);
	}
}