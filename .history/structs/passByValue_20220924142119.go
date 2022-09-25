package main

import "fmt"

type Point struct {
	X, Y float64
}

func main() {
	// Go 中的结构类型值是
	p1 := Point{1, 2}
	p2 := p1 // Copy of p1 is assigned to p2

	p2.X = 2

	fmt.Println(p1) // Output: {1 2}
	fmt.Println(p2) // Output: {2 2}
}