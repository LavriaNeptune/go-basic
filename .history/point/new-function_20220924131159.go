package main

import "fmt"

func main() {
	p := new(int)
	// new(T) 函数qu
	*p = 996

	fmt.Println("value", *p)
	fmt.Println("address", p)
}