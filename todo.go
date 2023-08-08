package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	ID     int
	Task   string
	Closed bool
}

type TodoList struct {
	Todos []Todo
}

func main() {
	fmt.Println("Todo List Manager")
	fmt.Println("Usage:")
	fmt.Println("add <task> - Add a new todo")
	fmt.Println("show - Show all todos")
	fmt.Println("edit <id> <task> - Edit a todo")
	fmt.Println("close <id> - Close a todo")

	todos := TodoList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		command := scanner.Text()

		parts := strings.Split(command, " ")
		switch parts[0] {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Invalid command. Usage: add <task>")
				continue
			}
			task := strings.Join(parts[1:], " ")
			todo := Todo{ID: len(todos.Todos) + 1, Task: task, Closed: false}
			todos.Todos = append(todos.Todos, todo)
			fmt.Printf("Added todo with ID %d\n", todo.ID)

		case "show":
			fmt.Println("Todo List:")
			if len(todos.Todos) == 0 {
				fmt.Println("No todos found.")
			} else {
				for _, todo := range todos.Todos {
					status := "Open"
					if todo.Closed {
						status = "Closed"
					}
					fmt.Printf("%d: [%s] %s\n", todo.ID, status, todo.Task)
				}
			}

		case "edit":
			if len(parts) < 3 {
				fmt.Println("Invalid command. Usage: edit <id> <task>")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid todo ID")
				continue
			}
			task := strings.Join(parts[2:], " ")
			for i := range todos.Todos {
				if todos.Todos[i].ID == id {
					todos.Todos[i].Task = task
					fmt.Printf("Edited todo with ID %d\n", id)
					break
				}
			}

		case "close":
			if len(parts) < 2 {
				fmt.Println("Invalid command. Usage: close <id>")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid todo ID")
				continue
			}
			for i := range todos.Todos {
				if todos.Todos[i].ID == id {
					todos.Todos[i].Closed = true
					fmt.Printf("Closed todo with ID %d\n", id)
					break
				}
			}

		default:
			fmt.Println("Invalid command")
		}
	}
}
