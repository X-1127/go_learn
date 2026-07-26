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

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.height + 2*r.width
}

type circle struct{
	radius float64
}

func (c circle) area() float64{
	return 3.14*c.radius*c.radius
}

func (c circle) perimeter() float64{
	return 2*3.14*c.radius
}

func main() {
	square := rect{width: 5, height: 10}
	cir := circle{radius:1}
	fmt.Println(shape.area(square))
	fmt.Println((shape.perimeter(square)))
	fmt.Println(shape.area(cir))
	fmt.Println((shape.perimeter(cir)))
}
