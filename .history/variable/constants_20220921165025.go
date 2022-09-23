package main

import "fmt"

const foo string = "Hello"

func main() {
	// 
	foo = 1;
	// 尝试修改将会报错:cannot assign to foo
    fmt.Println(foo)
}