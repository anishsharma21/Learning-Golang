package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func fetchurl() {
	start := time.Now()
	fmt.Printf("Time\tNumber of bytes\tURL\n")
    for _, url := range os.Args[1:] {
		startFetch := time.Now()
		hasHttpPrefix := strings.HasPrefix(url, "http://")
		if !hasHttpPrefix {
			url = "http://" + url
		}
        resp, err := http.Get(url)
        if err != nil {
            logError("fetchurl: reading %s: %v", url, err)
            continue
        }
		nbytes, err := io.Copy(io.Discard, resp.Body)
		if err != nil {
			log.Fatalf("fetchurl: reading %s: %v\n", url, err)
		}
        resp.Body.Close()
		secs := time.Since(startFetch).Seconds()
		fmt.Printf("%.2fs\t%7d\t%s\n", secs, nbytes, url)
    }
	fmt.Printf("\nfetch: %.2fs elapsed\n", time.Since(start).Seconds())
}

func logError(format string, v ...interface{}) {
    _, file, line, _ := runtime.Caller(1)
    parts := strings.Split(filepath.ToSlash(file), "/")
    shortFile := strings.Join(parts[len(parts)-2:], "/")
    log.Printf("%s:%d:%s", shortFile, line, fmt.Sprintf(format, v...))
}