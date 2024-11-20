package main

import (
	"fmt"
	"time"
)

func main() {
	printall()
	printall_c()
}

func printall() {
	for i := 0; i < 5 ; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Line %d\n", i+1)
	}
}

func printall_c() {
	ch := make(chan string)
	for i := 0; i < 5; i++ {
		go func(i int) {
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("Line %d\n", i+1)
		}(i)
	}

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Print(<-ch)
		}
		close(ch)
	}()

	time.Sleep(time.Second * 2)
}