package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// OLD code beneath showing just 1 goroutine
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	indexIter(1)
	// }()

	// how about 5 goroutines at the same time?
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go indexIter(i, &wg)
	}
	// this isn't a very common strat, normally you would use channels I think

	wg.Wait()
	fmt.Println("Done!")
}

func indexIter(gId int, pWg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		fmt.Printf("Goroutine %d: %d\n", gId, i)
	}
	pWg.Done()
}