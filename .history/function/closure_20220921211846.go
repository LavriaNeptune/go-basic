package main

import "fmt"

func myFunction() func(){
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

func myFunction() func(int) int {
	sum := 0

	return func(v int) int {
		sum += v

		return sum
	}
}