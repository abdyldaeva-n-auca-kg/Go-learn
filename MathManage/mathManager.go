package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hello, you are in the math manager system")
	fmt.Println("Please, print app you want access to(to exit print 'exit'): ")
	for scanner.Scan() {
		app := scanner.Text()
		switch app {
		case "salary manager":
			fmt.Println("You are in salary manager.")
			var salaries []int64
			for {
				var salary string
				fmt.Println("Enter salary(enter 'done' to finish): ")
				_, err := fmt.Scan(&salary)
				if err != nil {
					fmt.Println("Invalid input. Please enter valid integer.")
				}
				if salary == "done" {
					break
				} else {
					sal, err := strconv.ParseInt(salary, 10, 0)
					if err != nil {
						fmt.Println("Invalid input. Please enter valid integer.")
						continue
					}
					salaries = append(salaries, sal)
				}
			}
			fmt.Println(salaries)
			for {
				fmt.Println("Enter a command(sum, aver, min, max) or 'exit' to quit: ")
				var command string
				_, _ = fmt.Scan(&command)
				/*if err != nil {
					fmt.Println("Incorrect. Please, try again: ")
					continue
				}*/
				switch command {
				case "sum":
					var sum int64 = 0
					for _, salary := range salaries {
						sum += salary
					}
					fmt.Println("Sum of salaries:", sum)
				case "aver":
					sum := 0
					for _, salary := range salaries {
						sum += int(salary)
					}
					aver := sum / len(salaries)
					fmt.Println("Average of salary:", aver)
				case "min":
					min := salaries[0]
					for _, salary := range salaries {
						if salary < min {
							min = salary
						}
					}
					fmt.Println("Minimum salary:", min)
				case "max":
					max := salaries[0]
					for _, salary := range salaries {
						if salary > max {
							max = salary
						}
					}
					fmt.Println("Maximum salary:", max)
				case "exit":
					return
				default:
					fmt.Println("Incorrect command. Please try again.")
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
