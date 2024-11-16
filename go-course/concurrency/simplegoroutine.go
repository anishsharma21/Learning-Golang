package main

import (
	"fmt"
	"sync"
	"time"
)

func simplegoroutine() {
	var wg sync.WaitGroup;
	wg.Add(1)

	go func() {
		printMessage("Simple Goroutine");
		wg.Done();
	}();

	fmt.Println("Main function");
	wg.Wait();
}

func printMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message);
		time.Sleep(time.Millisecond * 500);
	}
}