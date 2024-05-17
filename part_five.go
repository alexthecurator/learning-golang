package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part_five() {
	fmt.Print("*** Calculator ***")
	fmt.Print("\n\nEnter values below Eg: (1 + 2 - 3 * 4 / 5 ...): \n")

	reader := bufio.NewScanner(os.Stdin)

	reader.Scan()
	values := strings.Split(reader.Text(), " ")

	var ans float64 = 0.0

	oper, _ := regexp.Compile(`[+\-*/]`)

	for i, text := range values {
		match := oper.MatchString(text)

		if match {
			b, _ := strconv.ParseFloat(values[i+1], 64)
			ans = arithmetics(text, ans, b)

			if i <= 2 {
				a, _ := strconv.ParseFloat(values[i-1], 64)
				ans = arithmetics(text, a, b)
			}
		}
	}

	fmt.Print("\nResult: ", ans)

}
