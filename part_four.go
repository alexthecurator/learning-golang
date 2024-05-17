package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func part_four() {
	// Todo: Convert to terminal user input
	sqrt := arithmetics("^", 64)
	sum := arithmetics("+", 3, 9, 10)
	minus := arithmetics("-", 1, 98, 100)
	divide := arithmetics("/", 3, 12) // Add 0 on the first param to simulate division be zero error
	multiply := arithmetics("*", 12, 12)

	fmt.Println("Sum:", sum, "| Minus:", minus, "| Multiply:", multiply, "| Divide:", divide, "| Sqrt:", sqrt)
}

func arithmetics(oper string, values ...float64) float64 {
	var result float64 = 0

	switch oper {
	case "*", "/", "**":
		result = 1
	}

	for _, n := range values {
		switch oper {
		case "+":
			result += n
		case "-":
			result = n - result
		case "*":
			result *= n
		case "/":
			if n == 0 {
				err := errors.New(": Cannot divide by zero")
				log.Fatal(err)
				return 0
			}
			result = n / result
		case "**":
			result = math.Pow(n, result)
		case "log":
			result = math.Log(n)
		}

	}

	return result
}
