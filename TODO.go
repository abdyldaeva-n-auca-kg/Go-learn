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
	package main

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "os"
)

type Task struct {
 Description string json:"description"
}

func loadTasksFromJSON(filename string) ([]Task, error) {
 tasks := []Task{}

 data, err := ioutil.ReadFile(filename)
 if err != nil {
  if os.IsNotExist(err) {
   return tasks, nil
  }
  return nil, err
 }

 err = json.Unmarshal(data, &tasks)
 if err != nil {
  return nil, err
 }

 return tasks, nil
}

func saveTasksToJSON(filename string, tasks []Task) error {
 data, err := json.Marshal(tasks)
 if err != nil {
  return err
 }

 err = ioutil.WriteFile(filename, data, 0644)
 if err != nil {
  return err
 }

 return nil
}

func addTask(taskDescription string, tasks *[]Task) {
 newTask := Task{Description: taskDescription}
 *tasks = append(*tasks, newTask)
}

func removeTask(taskDescription string, tasks *[]Task) {
 for i, task := range *tasks {
  if task.Description == taskDescription {
   *tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
   return
  }
 }
}

func displayTasks(tasks []Task) {
 fmt.Println("Todo List:")
 for i, task := range tasks {
  fmt.Printf("%d. %s\n", i+1, task.Description)
 }
}

func main() {
 const filename = "todo_list.json"
 tasks, err := loadTasksFromJSON(filename)
 if err != nil {
  fmt.Printf("Error loading tasks: %s\n", err)
  return
 }

 for {
  var choice int
  fmt.Println("\n1. Add Task\n2. Remove Task\n3. Display Tasks\n4. Exit")
  fmt.Print("Enter your choice (1/2/3/4): ")
  _, err := fmt.Scan(&choice)
  if err != nil {
   fmt.Println("Invalid input. Please try again.")
   continue
  }

  switch choice {
  case 1:
   var taskDescription string
   fmt.Print("Enter the new task: ")
   fmt.Scan(&taskDescription)
   addTask(taskDescription, &tasks)
   err = saveTasksToJSON(filename, tasks)
   if err != nil {
    fmt.Printf("Error saving tasks: %s\n", err)
   }
  case 2:
   var taskDescription string
   fmt.Print("Enter the task to remove: ")
   fmt.Scan(&taskDescription)
   removeTask(taskDescription, &tasks)
   err = saveTasksToJSON(filename, tasks)
   if err != nil {
    fmt.Printf("Error saving tasks: %s\n", err)
   }
  case 3:
   displayTasks(tasks)
  case 4:
   fmt.Println("Exiting the program.")
   return
  default:
   fmt.Println("Invalid choice. Please try again.")
  }
 }
}
}
