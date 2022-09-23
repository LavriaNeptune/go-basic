package main

import "fmt"



func main() {
	i := 42
	f := float64(i)
	u := uint(f)
	// 通过
	fmt.Printf("%T %T", f, u)
}