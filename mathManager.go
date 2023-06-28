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
			var salaries []int
			for {
				var salary int
				fmt.Println("Enter salaries(enter 'done' to finish): ")
				_, err := fmt.Scan(&salary)
				if err != nil {
					fmt.Println("Invalid input. Please enter valid integer.")
					continue
				}
				salaries = append(salaries, salary)

			}

			for {
				fmt.Println("Enter a command(sum, aver, min, max) or 'exit' to quit: ")
				var command string
				_, err := fmt.Scan(&command)
				if err != nil {
					fmt.Println("Incorrect. Please, try again: ")
					continue
				}
				switch command {
				case "sum":
					sum := 0
					for _, salary := range salaries {
						sum += salary
					}
					fmt.Println("Sum of salaries: %d\n", sum)
				case "aver":
					sum := 0
					for _, salary := range salaries {
						sum += salary
					}
					aver := sum / int(len(salaries))
					fmt.Println("Average of salary: %d\n", aver)
				case "min":
					min := salaries[0]
					for _, salary := range salaries {
						if salary < min {
							min = salary
						}
					}
					fmt.Println("Minimum salary: %d\n", min)
				case "max":
					max := salaries[0]
					for _, salary := range salaries {
						if salary > max {
							max = salary
						}
					}
					fmt.Println("Maximum salary: %d\n", max)
				case "exit":
					return
				default:
					fmt.Println("Invalid command. Please try again.")
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
		case "exit":
			os.Exit(0)
		}
	}
	if scanner.Err() != nil {
		fmt.Println("Ошибка чтения: ", scanner.Err())
	}
}
