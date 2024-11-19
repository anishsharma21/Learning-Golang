package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup() error {
	file, err := os.Open("dup.txt")
	if err != nil {
		return fmt.Errorf("failed opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCounts := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		lineCounts[line]++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	for line, count := range lineCounts {
		fmt.Printf("%d\t\"%s\"\n", count, line)
	}

	return nil
}