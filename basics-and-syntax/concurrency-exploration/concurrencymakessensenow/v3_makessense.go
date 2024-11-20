package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)
	var firstWg sync.WaitGroup
	firstWg.Add(1)

	go countV3(-1, c, &firstWg)
	go func() {
		for res := range c {
			fmt.Print(res)
		}
	}()

	firstWg.Wait()
	fmt.Printf("First wait group finished!\n\n")

	// now lets try with 5 goroutines
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go countV3(i, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for res := range c {
		fmt.Print(res)
	}
}

func countV3(grId int, c chan<- string, pWg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		c <- fmt.Sprintf("Goroutine %d: %d\n", grId, i)
	}
	if pWg != nil {
		pWg.Done()
	}
}