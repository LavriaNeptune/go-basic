package main

import "fmt"

func myFunction() func(int) int{
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
	// 1. Go 的函数
}

func main() {
	add := myFunction()
	add(5)
	fmt.Println(add(10))
}