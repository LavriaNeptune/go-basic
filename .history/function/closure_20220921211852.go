package main

import "fmt"

func myFunction() func(int) int{
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	add := myFunction()
	add(5)
	fmt.Println(add(10))
}
