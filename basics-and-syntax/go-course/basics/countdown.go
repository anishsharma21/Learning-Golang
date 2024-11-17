package main

import "fmt"

func countdown() {
	for count := 10; count > 0; count-- {
		fmt.Println(count);
	}
	fmt.Println("Liftoff!");
}