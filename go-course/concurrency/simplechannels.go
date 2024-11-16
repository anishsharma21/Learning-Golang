package main

import "fmt"

func main() {
	ch := make(chan int)

	go sendData(ch);

	for value := range ch {
		fmt.Println("Recieved:", value);
	}
}

func sendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i;
	}
	close(ch);
}