package main

import "fmt"

func myFunction() func(int) int{
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
	// 1. Go 的函数也可以作为值传递
	// 2. 
}

func main() {
	add := myFunction()
	add(5)
	fmt.Println(add(10))
}