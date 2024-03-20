package main

import (
	"fmt"
	"math"
)

func part_four() {
	sqrt := operations("^", 64)
	sum := operations("+", 3, 9, 10)
	minus := operations("-", 1, 98, 100)
	divide := operations("/", 2, 12)
	multiply := operations("*", 12, 12)

	fmt.Println("Sum:", sum, "| Minus:", minus, "| Multiply:", multiply, "| Divide:", divide, "| Sqrt:", sqrt)
}

func operations(oper string, values ...float64) float64 {
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
				result = n / result
			}

		}

		return result
	} else {
		result = math.Sqrt(values[0])
	}

	return result
}
