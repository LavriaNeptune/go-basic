package main

import "fmt"

var foo string = "Go is awesome"

// 可以只声明、不复制 -> var foo string

// 可以同时声明多个变量，多个变量需要使用小括号包裹起来
/* var (foo string = "Hello" 
		bar string ="World"
		baz int = 10086) */

// 声明时可以省略数据类型 -> 类型将由传入变量的值自动推断得出
// var foo = "What's my type?"

func main() {
    fmt.Println(foo)
}
