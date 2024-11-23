package main

import (
	"fmt"
	"time"
)

func main() {
	polling_select_exercise4()
}

func select_multichannels_exercise5() {

}

func polling_select_exercise4() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Message from channel 1"
		close(ch)
	}()

	loop: for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
			break loop
		case <-time.After(500 * time.Millisecond):
			fmt.Println("No data yet...")
		}
	}

	fmt.Println("Done!")
}