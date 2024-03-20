package main

import (
	"fmt"
	"reflect"
)

func part_two() {
	var number, amount uint8 = 10, 30

	x := "string"
	y := false

	fmt.Println(number, amount, x, y, "| Types:", reflect.TypeOf(number), reflect.TypeOf(amount), reflect.TypeOf(x), reflect.TypeOf(y))
}
