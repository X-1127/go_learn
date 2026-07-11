package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1 int = 19
	var b string = strconv.FormatInt(int64(n1), 10)
	fmt.Printf("b type:%T,value:%v", b, b)
}
