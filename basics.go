package main

import (
	"fmt"
	"reflect"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	part_one()
}

func part_two() {

}

func part_one() {
	emoji.Println("\nCheers to learning Golang :beer:")
	fmt.Println("Data types: ", reflect.TypeOf(102), reflect.TypeOf("..."), reflect.TypeOf(3.1423), reflect.TypeOf(false))
}
