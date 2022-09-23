package main

import "fmt"

func myFunction(values ...) int{
	
}

func main() {
	add := myFunction()
	add(5)
	fmt.Println(add(10)) // 15
}