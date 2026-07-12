package main

import "fmt"

func main() {
	fmt.Println("start or not:(input 1/0)")
	var start int
	fmt.Scanln(&start)
	if start == 0 {
		fmt.Println("end")
	} else {
		var sum int = 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		switch {
		case sum == 5050:
			fmt.Printf("correct result is %v", sum)
		case sum != 5050:
			fmt.Printf("error please check the code")
		}
	}
}
