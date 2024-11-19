package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo2() {
	start := time.Now().UnixMicro()
	fmt.Println(strings.Join(os.Args[1:], " "))
	totalTime := time.Now().UnixMicro() - start
	fmt.Printf("Total time (microseconds): %v\n", totalTime)
}