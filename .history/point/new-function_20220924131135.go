package main

import "fmt"

func main() {
	p := new(int)
	*p = 996

	fmt.Println("value", *p)
	fmt.Println("address", p)
}