package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu sync.Mutex
	value int
}

func (sc *SafeCounter) safeSet(num int) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.value = num
}

func main() {
	var sc SafeCounter
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sc.safeSet(i)
		}(i)
	}

	wg.Wait()
	fmt.Println("Final value:", sc.value)
}