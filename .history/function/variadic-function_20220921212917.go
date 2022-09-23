package main

import "fmt"

func myFunction(values ...int) int{
	sum :=0
	for _,i := range values {
		sum += i
	}
	return sum
}

func main() {
	sum
}