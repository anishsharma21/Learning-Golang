package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan string)
	var firstWg sync.WaitGroup

	firstWg.Add(1)
	go func() {
		countV3(0, ch1, &firstWg)
		close(ch1)
	}()

	for res := range ch1 {
		fmt.Print(res)
	}

	fmt.Println("First wait group done!")

	ch2 := make(chan string)
	var secondWg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		secondWg.Add(1)
		go countV3(i, ch2, &secondWg)
	}

	go func() {
		secondWg.Wait()
		close(ch2)
	}()

	for res := range ch2 {
		fmt.Print(res)
	}

	fmt.Println("Second wait group done!")
}

func countV3(grId int, c chan<- string, pWg *sync.WaitGroup) {
	if pWg != nil {
		defer pWg.Done()
	}
	for i := 1; i <= 5; i++ {
		c <- fmt.Sprintf("Goroutine %d: %d\n", grId, i)
	}
}