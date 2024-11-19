package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	Priority string `json:"priority"`
	DueDate time.Time `json:"dueDate"`
	Status string `json:"status"`
}

type TaskList []Task

var incr_id int = 0
const filename string = "tasks.json"
var taskList TaskList = TaskList{}

func main() {
	// check if file exists, create if it doesn't
	if !fileExists() {
		fmt.Println("tasks.json does not exist...")
		err := createFile()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("tasks.json file created!")
	}

	// load in data
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&taskList); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// parse and handle args
	err = parseArgs(os.Args)
	if err != nil {
		fmt.Println("Parsing error:", err)
	}
}

func parseArgs(args []string) error {
	if len(os.Args) <= 1 {
		return fmt.Errorf("no arguments provided")
	}

	// TODO complete switch statement to handle all commands
	switch args[1] {
	case "add":
		err := handleAdd(args)
		if err != nil {
			return err
		}
	// case "view":
	// 	handleView(args)
	// case "complete":
	// 	handleComplete(args)
	// case "remove":
	// 	handleRemove(args)
	// case "edit":
	// 	handleEdit(args)
	case "help":
		err := handleHelp(args)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid operation \"%s\"", args[1])
	}

	return nil
}

func handleAdd(args []string) error {
	// TODO consider unhappy case too
	var description, priority string
	var due_date time.Time

	if len(args) >= 3 {
		description = args[2]
	}

	for i := 5; i <= len(args) && i <= 7; i += 2 {
		err := parseAddFlag(3, args, &priority, &due_date)
		if err != nil {
			return err
		}
	}

	fmt.Println(description, priority, due_date)

	newTask := Task{ ID: generateId(), Description: description, Priority: priority, DueDate: due_date, Status: "pending" }
	taskList = append(taskList, newTask)
	return nil
}

func parseAddFlag(flagIndex int, args []string, priority *string, due_date *time.Time) error {
	if args[flagIndex] == "-priority" {
		if args[flagIndex + 1] != "low" && args[flagIndex + 1] != "medium" && args[flagIndex + 1] != "high" {
			return fmt.Errorf("invalid status \"%s\"", args[flagIndex + 1])
		}
		*priority = args[flagIndex + 1]
	} else if args[flagIndex] == "-due_date" {
		var err error
		*due_date, err = time.Parse("2006-01-02", args[flagIndex + 1])	
		if err != nil {
			return fmt.Errorf("invalid date format \"%s\", expected YYYY-MM-DD", args[flagIndex + 1])
		}
	}
	return nil
}

func handleHelp(args []string) error {
	if len(args) != 2 || args[1] != "help" {
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

func generateId() int {
	incr_id += 1
	return incr_id
}