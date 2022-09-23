package main

import "fmt"

func addTwoNumber(n1, n2 int) int {
	return n1 + n2
	
}

// 通过 func 关键字以如上格式声明函数 -> 所声明的函数可以在其他函数中进行调用...

func main() {
	result := addTwoNumber(1, 2)
	fmt.Println(result)
}