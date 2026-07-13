package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cal "pro5/package"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input the first int number:")
	var a, b int
	fmt.Scanln(&a)

	fmt.Println("Please input the second int number (not 0):")
	for {
		fmt.Scanln(&b)
		if b != 0 {
			break
		}
		fmt.Print("Second number cannot be 0, please re-enter: ")
	}

	fmt.Println("Please choose your cal method (add/sub/mul/div):")
	method, _ := reader.ReadString('\n')
	method = strings.TrimSpace(method)

	switch method {
	case "add":
		fmt.Println(cal.Add(a, b))
	case "sub":
		fmt.Println(cal.Sub(a, b))
	case "mul":
		fmt.Println(cal.Mul(a, b))
	case "div":
		result, err := cal.Div(a, b)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(result)
		}
	default:
		fmt.Println("wrong method spell")
	}
}
