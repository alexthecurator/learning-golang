package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part_five() {
	fmt.Println("[ Go Calculator 🧮 ] Ref this Eg: 1 + 2 - 3 * 4 / 5 | log n | ans + n etc..")
	fmt.Print("\n\nYour values (use ctrl + c to close): ")

	reader := bufio.NewScanner(os.Stdin)
	var ans float64 = 0.0
	var a float64 = 0.0
	var b float64 = 0.0

	for reader.Scan() {
		values := strings.Split(reader.Text(), " ")

		oper, _ := regexp.Compile(`[+\-*/log]`)

		for i, value := range values {
			if oper.MatchString(value) {

				// If answer is after operation
				// Assign the b to the previous answer or the next float value
				switch values[i+1] {
				case "ans":
					b = ans
				default:
					b = str_to_float(values[i+1])
				}

				if i == 0 {
					ans = arithmetics(value, b)
				}

				// Initial
				if i == 1 {
					// If answer is before operation
					// Assign the a to the previous answer or the first float value
					switch values[i-1] {
					case "ans":
						ans = arithmetics(value, b, ans)
					default:
						a = str_to_float(values[i-1])
						ans = arithmetics(value, b, a)
					}
				}

				// Then
				if i > 2 {
					ans = arithmetics(value, b, ans)
				}

			}
		}

		fmt.Println("\nAnswer (ans):", ans)
		fmt.Print("\nYour values (use ctrl + c to close): ")
	}
}

func str_to_float(value string) float64 {
	b, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal(err.Error(), " value:", b)
	}

	return b
}
