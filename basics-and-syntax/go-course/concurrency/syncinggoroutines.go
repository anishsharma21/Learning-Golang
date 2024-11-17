package main

import (
	"fmt"
	"sync"
)

func syncinggoroutines() {
	var wg sync.WaitGroup;

	for i := 1; i <= 3; i++ {
		wg.Add(1);
		go printNumbers(i, &wg);
	}

	wg.Wait();
	fmt.Println();
	fmt.Println("All goroutines finished.");
}

func printNumbers(id int, wg *sync.WaitGroup) {
	defer wg.Done();
	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i);
	}
}