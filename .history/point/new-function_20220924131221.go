package main

import "fmt"

func main() {
	p := new(int)
	// new(T) 函数取一个类型作为参数,返回参数类型的指针
	*p = 996

	fmt.Println("value", *p)
	fmt.Println("address", p)
}