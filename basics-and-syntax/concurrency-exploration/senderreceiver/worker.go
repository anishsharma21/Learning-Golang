package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 0; i < 10; i++ {
		go worker(jobs, results)
	}

	start := time.Now()
	for i := 0; i < 45; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 45; i++ {
		fmt.Printf("Fib %d: %d\n", i+1, <-results)
	}
	fmt.Printf("\n%.2fs elapsed\n", time.Since(start).Seconds())
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}