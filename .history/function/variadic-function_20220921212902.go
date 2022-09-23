package main

import "fmt"

func myFunction(values ...) int{
	sum :=0
	for _,i := range values {
		sum += i
	}
	return result
}

func main() {
	add := myFunction()
	add(5)
	fmt.Println(add(10)) // 15
}