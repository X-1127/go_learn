package cal

import "fmt"

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func Sub(num1 int, num2 int) int {
	return num1 - num2
}

func Mul(num1 int, num2 int) int {
	return num1 * num2
}

func Div(num1 int, num2 int) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return float64(num1) / float64(num2), nil
}
