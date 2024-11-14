package main

import (
	"fmt"
	"example/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go());
	fmt.Printf("Hello world, here is a number %d\n", 5);
	message := greetings.Hello("Anish")
	fmt.Println(message)
}
