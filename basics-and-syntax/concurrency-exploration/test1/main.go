package main

import (
	"fmt"
	"sync"
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
	var wg sync.WaitGroup
	ch := make(chan string)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("Line %d\n", i+1)
		}(i)
	}

	go func() {
        wg.Wait()
        close(ch)
    }()

	for msg := range ch {
		fmt.Print(msg)
	}
}