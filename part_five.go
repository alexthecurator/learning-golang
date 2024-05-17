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
	fmt.Print("[ Your Go Calculator 🧮 ] \nEnter values below Eg: 1 + 2 - 3 * 4 / 5 etc.. :\n")

	reader := bufio.NewScanner(os.Stdin)

	reader.Scan()
	values := strings.Split(reader.Text(), " ")

	var ans float64 = 0.0

	oper, _ := regexp.Compile(`[+\-*/]`)

	for i, value := range values {
		if oper.MatchString(value) {
			b, err := strconv.ParseFloat(values[i+1], 64)
			if err != nil {
				log.Fatal(err.Error())
			}

			// Initially
			if i == 1 {
				a, err := strconv.ParseFloat(values[i-1], 64)
				if err != nil {
					log.Fatal(err.Error())
				}
				ans = arithmetics(value, b, a)
			}

			// After
			if i > 2 {
				ans = arithmetics(value, b, ans)
			}
		}
	}

	fmt.Println("\nResult:", ans)

}
