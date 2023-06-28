package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hello, you are in the math manager system")
	fmt.Println("Please, print app you want to access(to exit print 'exit'): ")
	for scanner.Scan() {
		app := scanner.Text()
		switch app {
		case "salary manager":
			fmt.Println("You are in salary manager.")
			var salaries []float64
			for {
				var salary float64
				fmt.Println("Enter salary(enter 'done' to finish): ")
				_, err := fmt.Scanf("%f", &salary)
				if err != nil {
					fmt.Println("Invalid input. Please enter valid integer.")
					continue
				}
				salaries = append(salaries, salary)
			}

			for {
				fmt.Println("Enter a command(sum, aver, min, max) or 'exit' to quit: ")
				var command string
				_, err := fmt.Scanln(&command)
				if err != nil {
					fmt.Println("Incorrect. Please, try again: ")
					continue
				}
				switch command {
				case "sum":
				case "aver":
				case "min":
				case "max":
				}
			}
		case "calculator":
			fmt.Println("You are in calculator. Print command: ")
			for scanner.Scan() {
				command := scanner.Text()
				switch command {
				case "sum":
					fmt.Println("Print two integers")

				}
			}
			if app == "exit" {
				break
			}

		}
	}
	if scanner.Err() != nil {
		fmt.Println("Ошибка чтения: ", scanner.Err())
	}
}
