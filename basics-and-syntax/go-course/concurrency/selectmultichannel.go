package main

import "fmt"

func multichannel() {
	ch1 := make(chan string);
	ch2 := make(chan string);

	go func() {
		ch1 <- "Message for ch1";
	}();

	go func() {
		ch2 <- "Message for ch2";
	}();

	msg1 := <-ch1;
	fmt.Println(msg1);

	msg2 := <-ch2
	fmt.Println(msg2);
}