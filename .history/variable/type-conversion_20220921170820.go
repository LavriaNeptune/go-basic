package main

import "fmt"



func main() {
	i := 42
	f := float64(i)
	u := uint(f)
	// ιθΏε½ε
	fmt.Printf("%T %T", f, u)
}