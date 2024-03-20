package main

import (
	"fmt"
)

func part_three() {
	a := sumA(13, 12)    // * short hand var declaration
	var b = sumB(13, 12) // * full declaration

	first, second, third := primitives() // * 🚀 destructuring

	fmt.Println("Sum of A =", a, ", Sum of B =", b)
	fmt.Println("Primitives:", first, second, third)
	variadic(20, 32, 49, 120, 1000)
}

/*
* This Function has declaritive return type written as:
* func name (params..) uint8 { ..body }
 */
func sumA(a uint8, b uint8) uint8 {
	sum := a + b

	return sum
}

/*
* This Function has implicit return written as:
* func name (..params) (returning var name and type) { ..body return; }
 */
func sumB(a uint8, b uint8) (sum uint8) {
	sum = a + b
	return
}

/*
* This Function has implicit return written as:
* func name (..params) (returning types) { ..body return; }
 */
func primitives() (uint8, uint8, uint16) {
	return 100, 200, 300
}

func variadic(numbers ...int) {
	fmt.Println("Variadics:", numbers)
}
