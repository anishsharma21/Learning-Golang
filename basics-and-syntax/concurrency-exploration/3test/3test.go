package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("First function:", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("Second function:", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	wg.Wait()

	fmt.Println("Done with first part!")

	c := make(chan int)
	go func() {
		sum := 0
		for i := 0; i < 100; i++ {
			sum += i
		}
		c <- sum
	}()

	output := <-c
	fmt.Println("Total sum:", output)
}