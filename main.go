package main

import "fmt"

func main() {
	functions := []func(){part_one, part_two, part_three, part_four, part_five}

	for i, exec := range functions {
		fmt.Println("\nPart:", i+1)
		exec()
	}
}
