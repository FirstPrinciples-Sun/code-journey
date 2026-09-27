package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	ID   int
	Task string
}

var todos []Todo
var nextID int = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a command (add, list, delete, exit):")
		scanner.Scan()
		command := strings.TrimSpace(scanner.Text())

		switch command {
		case "add":
			fmt.Println("Enter a task:")
			scanner.Scan()
			task := strings.TrimSpace(scanner.Text())
			addTodo(task)
		case "list":
			listTodos()
		case "delete":
			fmt.Println("Enter the ID of the task to delete:")
			scanner.Scan()
			idStr := strings.TrimSpace(scanner.Text())
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			deleteTodo(id)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}

func addTodo(task string) {
	todo := Todo{ID: nextID, Task: task}
	todos = append(todos, todo)
	nextID++
	fmt.Println("Task added:", task)
}

func listTodos() {
	if len(todos) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, todo := range todos {
		fmt.Printf("%d: %s\n", todo.ID, todo.Task)
	}
}

func deleteTodo(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Task deleted:", todo.Task)
			return
		}
	}
	fmt.Println("Task not found with ID:", id)
}