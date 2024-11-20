package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func fetchall() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		hasHttpPrefix := strings.HasPrefix(url, "http://")
		if !hasHttpPrefix {
			url = "http://" + url
		}
		go fetchconcurrent(url, ch)
	}
	fmt.Printf("Time\tNumber of bytes\tURL\n")
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("\nfetchall: %.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchconcurrent(url string, ch chan<-string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", secs, nbytes, url)
}