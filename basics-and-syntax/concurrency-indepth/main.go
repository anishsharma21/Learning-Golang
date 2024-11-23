package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
}

func channel_done_results(numberOfWorkers int, numberOfJobs int, channelCapacity int) {
	jobs := make(chan int, channelCapacity)
	results := make(chan int, channelCapacity)
	done := make(chan bool, channelCapacity)

	for range numberOfWorkers {
		go worker_bufchannel_done_results(jobs, results)
	}

	start := time.Now()
	go func() {
		for i := 0; i < numberOfJobs; i++ {
			jobs <- i+1
		}
		close(jobs)
	}()

	go func() {
		for result := range results {
			fmt.Println("Job complete:", result)
			done <- true
		}
		close(done)
	}()

	for i := 0; i < numberOfJobs; i++ {
		<-done
	}
	close(results)
	fmt.Printf("Time elapsed: %dµs\n", time.Since(start).Microseconds())
}

func worker_bufchannel_done_results(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		time.Sleep(200 * time.Millisecond)
		results <- job * job
	}
}

// problem with the job queue below is that even if a job is complete, we can't print
// that it is finished until all the jobs are added to the job queue because it blocks
// so, to print when any single job is complete, we can create goroutines for that reason
// which receive from the done channel, so 2 sets of workers (idk how good this pattern is)
// FIX: you wrap the job creation logic in a goroutine
func done_worker_bufchannel_jobqueue(done <-chan int) {
	for jobDone := range done {
		fmt.Printf("Job complete: %v\n", jobDone)
	}
}

// this is sort of like a rate limiter where you can set the buffer for number of incoming
// jobs at any given moment, then create as many workers as you like to handle them
// important note: each point where channel receives or sends is potentially blocking if:
// the buffered channel is not full, which means nothing below it runs until its done
func bufchannel_jobqueue() {
	jobs := make(chan int, 20)
	done := make(chan int, 20)

	for range 4 {
		go bufchannel_jobqueue_worker(jobs, done)
	}

	start := time.Now()
	const numjobs int = 20

	for i := 0; i < numjobs; i++ {
		jobs <- i+1
	}
	close(jobs)

	for range numjobs {
		fmt.Printf("Job complete: %d\n", <-done)
	}
	close(done)

	fmt.Printf("%.3fs elapsed\n", time.Since(start).Seconds())
	fmt.Println("All jobs complete")
}

func bufchannel_jobqueue_worker(jobs <-chan int, done chan<- int) {
	for job := range jobs {
		fmt.Println("Processing job:", job)
		time.Sleep(time.Second)
		done <- job
	}
}

// btw, unbuffered channels block because they can't receive anything else until they are emptied
// by blocking, they allow goroutine to empty that single unbuffered channel
// buffered channels don't fill immediately, so they don't always block
// but when buffered channels are filled, they end up blocking the same way as unbuffered channels
func dup_channels_exercise4() {
	// fan-out pattern
	jobs := make(chan int)
	done := make(chan bool)

	for range 3 {
		go dup_worker_channels_exercise4(jobs, done)
	}

	for i := 0; i < 3; i++ {
		jobs <- i+1
	}
	close(jobs)

	for range 3 {
		<-done
	}
}

func dup_worker_channels_exercise4(jobs <-chan int, done chan<- bool) {
	for job := range jobs {
		fmt.Println("Doing work...")
		time.Sleep(1 * time.Second)
		fmt.Printf("Job %d done!\n", job)
	}
	done <- true
}

func channels_exercise4() {
	// fan-out pattern
	jobs := make(chan int)

	for i := 0; i < 3; i++ {
		go worker_channels_exercise4(jobs)
	}

	for i := 0; i < 3; i++ {
		jobs <- i+1
	}
	close(jobs)

	time.Sleep(4 * time.Second)
}

func worker_channels_exercise4(jobs <-chan int) {
	for job := range jobs {
		fmt.Println("Doing work...")
		time.Sleep(1 * time.Second)
		fmt.Printf("Job %d done!\n", job)
	}
}

func channels_exercise3() {
	// deliberate deadlock creation and commented out solution
	ch := make(chan int)

	ch <- 1

	// go func() {
	// 	ch <- 1
	// }()

	fmt.Println("This line is not reached... channel send is blocking")
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
