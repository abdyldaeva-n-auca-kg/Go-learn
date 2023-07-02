package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TODO struct {
	ID     int
	Task   string
	Closed bool
}

type TodoList struct {
	Todos []TODO
}

type Todos struct {
	ID     int
	Task   string
	Closed bool
}

func main() {
	fmt.Println("TODO list manager")

	fmt.Println("----------\n" +
		"| Usage: |\n" +
		"----------\n" +
		"add <task name> - Add a new task in todo list\n" +
		"show - Show all tasks\n" +
		"edit <id>- Edit task.(You can edit one at time and few/all also)\n" +
		"close <id>- Exits TODO\n" +
		"help - Display usage of the program")

	todos := TodoList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Terminal: ")
		scanner.Scan()
		command := scanner.Text()

		parts := strings.Split(command, " ")
		switch parts[0] {
		default:
			fmt.Println("Invalid command")
		case "add":
			if len(parts) < 2 {
				fmt.Println("Not enough arguments. Usage: add <task name>")
				continue
			}
			task := strings.Join(parts[1:], " ")
			todo := Todo{ID: len(todos.Todos) + 1, Task: task, Closed: false}
			todos.Todos = append(todos.Todos, todo)
			fmt.Println("Added task to the list with ID %d\n", todo.ID)
		case "show":
			if len(todos.Todos) == 0 {
				fmt.Println("TODO list is empty")
			} else {
				fmt.Println("Todo List:")
				for _, todo := range todos.Todos {
					status := "Open"
					if todo.Closed {
						status = "Closed"
					}
					fmt.Println("%d:[%s] %s\n", todo.ID, status, todo.Task)
				}
			}
		case "edit":
			if len(parts) < 3 {
				fmt.Println("Not enougn arguments. Usage: edit <id> <task>")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			task := strings.Join(parts[2:], " ")
			for i := range todos.Todos {
				if todos.Todos[i].ID == id {
					todos.Todos[i].Task = task
					fmt.Println("Edited task with ID %d\n", id)
					break
				}
			}
		case "close":
			if len(parts) < 2 {
				fmt.Println("Invalid command/Not enough arguments. Usage: close <id>")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid task ID")
				continue
			}
			for i := range todos.Todos {
				if todos.Todos[i].ID == id {
					todos.Todos[i].Closed = true
					fmt.Printf("Closed task with ID %d\n", id)
					break
				}
			}
		}
	}
}
