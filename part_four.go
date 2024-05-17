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
	divide := arithmetics("/", 0, 12)
	multiply := arithmetics("*", 12, 12)

	fmt.Println("Sum:", sum, "| Minus:", minus, "| Multiply:", multiply, "| Divide:", divide, "| Sqrt:", sqrt)
}

func arithmetics(oper string, values ...float64) float64 {
	var result float64 = 0
	if oper != "^" {
		switch oper {
		case "*", "/":
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
					err := errors.New(": Division by zero, is not allowed")
					log.Fatal(err)
					return 0
				}
				result = n / result
			}
		}

	} else {
		result = math.Sqrt(values[0])
	}

	return result
}
