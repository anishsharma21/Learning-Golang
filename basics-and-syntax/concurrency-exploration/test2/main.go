package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go boring("Boring!", ch)
	for i := 0; i < 5; i++ {
		fmt.Printf("You said: %q\n", <-ch)
	}
	fmt.Println("Done!")
}

func boring(msg string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Second)
	}
}