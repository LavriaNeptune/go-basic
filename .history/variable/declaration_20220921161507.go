package main

import "fmt"

var foo string = "Go is awesome"
// 可以只声明、不复制 -> var foo string
// 可以同时声明多个变量，多个变量需要使用小

var (foo string = "Hello" 
		bar string ="World"
		baz int = 10086)

func main() {
    fmt.Println(foo)
		fmt.Println(baz)
}
