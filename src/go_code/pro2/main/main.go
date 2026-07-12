package main

import (
	"fmt"
)

func main() {
	var n1 int = 19
	var b string = fmt.Sprintf("%d", n1)
	fmt.Printf("b type:%T,value:%v", b, b)
}
