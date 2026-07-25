package main

import (
	"fmt"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r *rect) area() float64 {
	r.width += 1
	return r.width * r.height
}
func (r *rect) perimeter() float64 {
	return 2*r.height + 2*r.width
}

func main() {
	square := rect{width: 5, height: 10}

	fmt.Println(shape.area(&square))
	fmt.Println((shape.perimeter(&square)))
}
