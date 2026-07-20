package main

import "fmt"

type class struct {
	teacher string
	student string
}

func main() {
	c := class{teacher: "1", student: "2"}

	fmt.Println(c.teacher)
}
