package main

import "fmt"

type Point struct {
	X, Y float64
}

func main() {
	// Go 中的结构类型值是按值传递的(别的语言如js 可能是按引用传递的)
	p1 := Point{1, 2}
	p2 := p1 // Copy of p1 is assigned to p2

	p2.X = 1000

	fmt.Println(p1) // Output: {1 2}
	fmt.Println(p2) // Output: {1000 2}
	// 可以看出修改不会影响到原值
}