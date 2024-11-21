package main

import (
	"fmt"
	"time"
)

func main() {
    first := time.Now()
    num := fib_memo(100)
    firstEnd := time.Since(first).Nanoseconds()
    fmt.Println("Fib num:", num)
    fmt.Printf("First: %.2f nanoseconds elapsed\n", float64(firstEnd))

    second := time.Now()
    num = fib_space(100)
    secondEnd := time.Since(second).Nanoseconds()
    fmt.Println("Fib num:", num)
    fmt.Printf("Second: %.2f nanoseconds elapsed\n", float64(secondEnd))

	// after global memoization, should run much faster
	// and it does, almost as fast as iterative approach
    third := time.Now()
    num = fib_memo(100)
    thirdEnd := time.Since(third).Nanoseconds()
    fmt.Println("Fib num:", num)
    fmt.Printf("Third: %.2f nanoseconds elapsed\n", float64(thirdEnd))

	// now lets see if we can make recursive approach even faster with worker pools
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)

	for i := 0; i < 45; i++ {
		jobs <- i
	}
	close(jobs)

	fourth := time.Now()
	for i := 0; i < 45; i++ {
		<-results
	}
	close(results)
    fourthEnd := time.Since(fourth).Nanoseconds()
    fmt.Printf("Fourth: %.2f nanoseconds elapsed\n", float64(fourthEnd))
}

func worker(jobs <-chan int, results chan<- int) {
	for val := range jobs {
		results <- fib_memo(val)
	}
}

var fibMemo = make(map[int]int)

func fib_memo(n int) int {
    v, ok := fibMemo[n]
    if ok {
        return v
    }

    if n <= 1 {
        return n
    }

    fibMemo[n] = fib_memo(n-1) + fib_memo(n-2)
    return fibMemo[n]
}

func fib_space(n int) int {
    if n <= 1 {
        return n
    }

    prev, curr := 0, 1
    for i := 2; i <= n; i++ {
        temp := curr
        curr += prev
        prev = temp
    }
    return curr
}