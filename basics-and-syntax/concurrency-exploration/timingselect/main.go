package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "Channel 1 message"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Channel 2 message"
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	for {
		select {
		case msg := <- ch1:
			fmt.Println(msg)
		case msg := <- ch2:
			fmt.Println(msg)
		}
	}
}