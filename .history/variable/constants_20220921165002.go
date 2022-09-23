package main

import "fmt"

const foo string = "Hello"

func main() {
	foo = 1;
	// 报错:cannot assign to foo
    fmt.Println(foo)
}