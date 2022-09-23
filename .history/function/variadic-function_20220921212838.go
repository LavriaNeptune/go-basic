package main

import "fmt"

func myFunction(values ...) int{
	sum :=0
	for _,i := range values {
		
	}
}

func main() {
	add := myFunction()
	add(5)
	fmt.Println(add(10)) // 15
}