package main

import "fmt"

func myFunction(values ...) func(int) int{
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
	// 1. Go 的函数也可以作为值传递
	// 2. 函数作为返回值 return 后,依旧能够访问其声明处作用域中的变量...这种现象也叫做闭包
}

func main() {
	add := myFunction()
	add(5)
	fmt.Println(add(10)) // 15
}