package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Task struct represents a task in the TODO list
type Task struct {
	ID   int
	Name string
}

// TodoList struct represents the list of tasks
type TodoList struct {
	Tasks []Task
}

// AddTask adds a new task to the list
func (list *TodoList) AddTask(name string) {
	taskID := len(list.Tasks) + 1
	task := Task{ID: taskID, Name: name}
	list.Tasks = append(list.Tasks, task)
	fmt.Println("Task added successfully!")
}

// RemoveTask removes a task from the list by its ID
func (list *TodoList) RemoveTask(id int) {
	for i, task := range list.Tasks {
		if task.ID == id {
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
			fmt.Println("Task removed successfully!")
			return
		}
	}
	fmt.Println("Task not found")
}

// PrintTasks prints all tasks in the list
func (list *TodoList) PrintTasks() {
	if len(list.Tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range list.Tasks {
		fmt.Printf("%d: %s\n", task.ID, task.Name)
	}
}

func main() {
	var list TodoList

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTODO List Menu:")
		fmt.Println("1. Add Task")
		fmt.Println("2. Remove Task")
		fmt.Println("3. View Tasks")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid choice. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter task name: ")
			scanner.Scan()
			taskName := scanner.Text()
			list.AddTask(taskName)
		case 2:
			fmt.Print("Enter task ID to remove: ")
			scanner.Scan()
			taskID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid task ID. Please enter a number.")
				continue
			}
			list.RemoveTask(taskID)
		case 3:
			list.PrintTasks()
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}
