package main

import "fmt"

func main() {
	fmt.Println("please input the value of n1:")
	var n1 int
	fmt.Printf("n1 location:%v\n", &n1)
	fmt.Scanln(&n1)
	var prt *int = &n1
	fmt.Printf("prt value:%v,*prt value:%v\n", prt, *prt)
	fmt.Scanln(prt)
	fmt.Printf("prt value:%v,*prt value:%v\n", prt, *prt)
}
