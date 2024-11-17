package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	id int
	description string
	priority string
	dueDate time.Time
	status string
}

type TaskList struct {
	tasks []Task
}

const filename string = "tasks.json"

func main() {
	if fileExists() {
		fmt.Println("File exists!")
	} else {
		fmt.Println("File does not exist...")
	}
}

func fileExists() bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}