// 被注释的代码为一个经典错误代码示例
package main

import "fmt"

// func exchange(num1 int, num2 int) {
// 	var t int
// 	t = num1
// 	num2 = num1
// 	num1 = t
// }

// func main() {
// 	var num1 int = 20
// 	var num2 int = 10
// 	fmt.Printf("before ex:num1 = %v,num2 = %v \n", num1, num2)
// 	exchange(num1, num2)
// 	fmt.Printf("after ex:num1 = %v,num2 = %v \n", num1, num2)
// }

func exchange(pa *int, pb *int) {
	var t int
	t = *pa
	*pa = *pb
	*pb = t
}

func main() {
	var num1 int = 20
	var num2 int = 10
	fmt.Printf("before ex:num1 = %v,num2 = %v \n", num1, num2)
	exchange(&num1, &num2)
	fmt.Printf("after ex:num1 = %v,num2 = %v \n", num1, num2)
}
