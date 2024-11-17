package main

import (
	"errors"
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
var taskList TaskList = TaskList{}
var commands = [5]string{ "add", "view", "remove", "edit", "help" }

func main() {
	if !fileExists() {
		fmt.Println("tasks.json does not exist...")
		err := createFile()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("tasks.json file created!")
	}
	err := parseArgs(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func parseArgs(args []string) error {
	// check for invalid command
	var validOperation bool = false
	for _, cmd := range commands {
		if args[1] == cmd {
			validOperation = true
			break
		}
	}
	if !validOperation {
		return fmt.Errorf("Invalid operation: %s", args[1])
	}
	// TODO should use a switch statement instead
	switch args[1] {
	case "add":
		handleAdd(args)
	case "view":
		handleView(args)
	case "remove":
		handleRemove(args)
	case "edit":
		handleEdit(args)
	case "help":
		err := handleHelp(args)
		if err != nil {
			return err
		}
	}

	return nil
}

func handleHelp(args []string) error {
	// check if generic help required
	if len(args) != 2 || args[1] == "help" {
		return errors.New("invalid arguments for help")
	}

    bold := "\033[1m"
    reset := "\033[0m"
    fmt.Println(bold + "tasklist add" + reset)
	fmt.Println("task_name\ta string representing the description of the task")
    fmt.Println("-priority\tcan be used to assign priority level (low, medium, high)")
    fmt.Println("-due_date\tcan be used to specify due date (YYYY-MM-DD)")
    fmt.Println("\nexample: tasklist add \"Finish homework\" -priority high -due_date 2024-11-18")
	// to be complete
	return nil
}

func createFile() error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return errors.New("error creating tasks.json file")
	}
	file.WriteString("[]")
	file.Close()
	return nil
}

func fileExists() bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}