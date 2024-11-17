package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Task struct {
	Id int
	Description string
}

type TaskList struct {
	Tasks []Task
}

var taskList TaskList = TaskList{}
var id int = 0

func main() {
	var userchoice string

	for {
		fmt.Println("Choices:")
		fmt.Println("v - View task list")
		fmt.Println("a - Add a task")
		fmt.Println("d - Delete a task")
		fmt.Println("x - Exit program")
		fmt.Print("\nChoice: ")
		fmt.Scan(&userchoice)

		switch userchoice {
		case "v":
			viewTaskList()
		case "a":
			err := addTask()
			if err != nil {
				fmt.Println("\n\033[31mError:", err)
				fmt.Print("\033[0m")
			}
		case "d":
			err := deleteTask()
			if err != nil {
				fmt.Println("\n\033[31mError:", err)
				fmt.Print("\033[0m")
			}
		case "x":
			return
		default:
			fmt.Println("\nInvalid choice. Choose from v, a, d, or x.")
		}
		fmt.Println()
	}
}

func viewTaskList() {
	fmt.Println("\nTask list:")

	if len(taskList.Tasks) == 0 {
		fmt.Println("[]")
		return
	}

	for i := 0; i < len(taskList.Tasks); i++ {
		fmt.Printf("[ID: %d] - %s", taskList.Tasks[i].Id, taskList.Tasks[i].Description)
	}
}

func addTask() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Task description: ")
	description, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading from stdin:", err)
		return errors.New("could not create task")
	}
	if len(description) <= 1 {
		return errors.New("task description cannot be empty")
	}
	newtask := Task{ Id: generateId(), Description: description }
	taskList.Tasks = append(taskList.Tasks, newtask)
	fmt.Println("\n\033[32mTask successfully added!\033[0m");
	viewTaskList();
	return nil
}

func deleteTask() error {
	if len(taskList.Tasks) == 0 {
		return errors.New("task list is empty")
	}
	var taskid int

	fmt.Print("Task's ID: ")
	_, err := fmt.Scan(&taskid)
	if err != nil {
		return err
	}

	var taskIndex int
	for i, v := range taskList.Tasks {
		if v.Id == taskid {
			taskIndex = i
			break
		}
		if i == len(taskList.Tasks) - 1 {
			return fmt.Errorf("task with ID = %d not found", taskid)
		}
	}

	taskList.Tasks = append(taskList.Tasks[:taskIndex], taskList.Tasks[taskIndex+1:]...)
	return nil
}

func generateId() int {
	id += 1
	return id
}