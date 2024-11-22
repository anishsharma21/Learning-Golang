package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	channels_exercise2()
}

func channels_exercise2() {
	ch := make(chan int, 3)

	for i := 0; i < cap(ch); i++ {
		ch <- i + 1
		fmt.Println("Number of items in buffered channel:", len(ch))
	}
	close(ch)

	for val := range ch {
		fmt.Println(val)
		fmt.Println("Number of items in buffered channel:", len(ch))
	}

	// for i := 0; i < cap(ch); i++ {
	// 	fmt.Println(<-ch)
	// 	fmt.Println("Number of items in buffered channel:", len(ch))
	// }
}

func channels_exercise1() {
	ch := make(chan string)

	go func() {
		ch <- "message from goroutine"
	}()

	fmt.Println(<-ch)
}

func eleventh_test() {
	done := make(chan bool)

	go worker_eighth_test(done)
	go worker_eleventh_test(done)

	// in this one, only 1 goroutine sends result back
	// because there is only 1 channel receive
	<-done
	fmt.Println("Main function exits")
}

func worker_eleventh_test(ch chan bool) {
	fmt.Println("Doing work in 11th test worker...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done work in 11th test worker!")
	ch <- true
}

func tenth_test() {
	// same as ninth test but no buffered channel
	done := make(chan bool)

	go worker_eighth_test(done)
	go worker_eighth_test(done)

	// works the same since both below are also blocking
	<-done
	<-done
	fmt.Println("Main function exits")
}

func ninth_test() {
	done := make(chan bool, 2)

	go worker_eighth_test(done)
	go worker_eighth_test(done)

	// both are blocking because buffered channel is full
	<-done
	<-done
	fmt.Println("Main function exits")
}

func eighth_test() {
	// synchronising goroutines with channels where goroutine does work
	done := make(chan bool)

	go worker_eighth_test(done)
	// blocked until channel receives done
	<-done
	fmt.Println("Main function exits")
}

func worker_eighth_test(done chan bool) {
	fmt.Println("Doing work...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done work!")
	done <- true
}

func seventh_test() {
	// buffered channels can hold values and don't block while its not full
	// so they don't need immediate receiving
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func sixth_test() {
	ch := make(chan string)

	go worker_sixth_test(ch)

	output := <-ch
	fmt.Println("Worker output:", output)
}

func worker_sixth_test(ch chan string) {
	ch <- "message from the worker"
}

func fifth_test() {
	ch := make(chan int)

	go func() {
		// sending blocks until main goroutine ready to receive
		ch <- 21
		// which means this line below may not always get called
		fmt.Println("Sending '21' into the channel")
	}()

	// receiving blocks until goroutine sends value into channel
	value := <-ch
	fmt.Println("Received value:", value)
}

func fourth_test() {
	for i := 1 ; i <= 10; i++ {
		go func() {
			fmt.Println(i*i)
			fmt.Println("Number of goroutines:", runtime.NumGoroutine())
		}()
	}
	time.Sleep(2 * time.Second)
}

func third_test() {
	// printing the number of goroutines running at any given moment
	go displayNumberOfGoroutines()
	go printLettersWithDelay()
	go printLettersWithDelay()
	go printNumbersWithDelay()
	go printNumbersWithDelay()

	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
}

func displayNumberOfGoroutines() {
	for {
		fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
		fmt.Println("Number of CPU cores avalable:", runtime.NumCPU())
		time.Sleep(500 * time.Millisecond)
	}
}

func second_test() {
	// letters and numbers are interleaved
	go printNumbersWithDelay()
	go printLettersWithDelay()

	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
}

func printNumbersWithDelay() {
	for i := 0;; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLettersWithDelay() {
	for ch := 'a';; ch++ {
		fmt.Printf("%c\n", ch)
		time.Sleep(150 * time.Millisecond)
	}
}
func first_test() {
	// go routine gets created, runs independently
	// goroutines can only be blocking if synchronisation methods are used
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Goroutine: %d\n", i)
		}
	}()
	// the Go runtime now needs to handle the main thread and the goroutine
	// since so many normal prints need to be made in the loop below, most of the time, all the
	// goroutine threads end up printing out as they are interleaved during runtime
	for i := 0; i < 100; i++ {
		fmt.Printf("Normal: %d\n", i)
	}
}