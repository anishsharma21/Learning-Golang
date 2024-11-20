package main

import "fmt"

func v1() {
	go count()
	count()
	fmt.Scanln()
	// this prints from goroutine too because scanln is blocking
}

func count() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}