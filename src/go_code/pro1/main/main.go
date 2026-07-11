package main //声明文件所在的包，每个go文件必须有归属的包

import (
	"fmt"
	helper "pro1/package"
) //引入程序中需要用的包，为了使用包中的函数，比如：Println

func main() { //主函数 程序的入口
	fmt.Println("Hello golang")
	fmt.Println(helper.Hello())
}
