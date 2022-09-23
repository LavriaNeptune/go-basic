package main

import "fmt"

const foo string = "Hello"

func main() {
	// 尝试修改将会报错
	foo = 1;
	// 报错:cannot assign to foo
    fmt.Println(foo)
}