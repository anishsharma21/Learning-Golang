package main

import (
	"fmt"
	"os"
	"time"
)

func echo1() {
	start := time.Now().UnixMicro()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s);
	totalTime := time.Now().UnixMicro() - start
	fmt.Printf("Total time (microseconds): %v\n", totalTime)
}