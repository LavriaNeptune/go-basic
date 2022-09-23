package main

import "fmt"

const foo string = "Hello"

func main() {
	// foo = 1;
	// xia尝试修改常量 foo 的值将会报错:cannot assign to foo
    fmt.Println(foo)
}