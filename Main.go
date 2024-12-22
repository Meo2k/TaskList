package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task structure
type Task struct {
	ID        int
	Title     string
	Completed bool
}

// Global task list
var tasks []Task
var nextID int = 1

// Add a new task
func addTask(title string) {
	task := Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Task added successfully!")
}

// List all tasks
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("\nTask List:")
	for _, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, status, task.Title)
	}
}

// Mark a task as completed
func completeTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Println("Task marked as completed!")
			return
		}
	}
	fmt.Println("Task not found!")
}

// Delete a task
func deleteTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			return
		}
	}
	fmt.Println("Task not found!")
}

// Main menu
func mainMenu() {
	fmt.Println("\nTask Manager")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Complete Task")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Exit")
	fmt.Print("Choose an option: ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		mainMenu()
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())
			addTask(title)
		case "2":
			listTasks()
		case "3":
			fmt.Print("Enter task ID to complete: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text()) // string to interger 
			if err != nil {
				fmt.Println("Invalid ID!")
			} else {
				completeTask(id)
			}
		case "4":
			fmt.Print("Enter task ID to delete: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid ID!")
			} else {
				deleteTask(id)
			}
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
