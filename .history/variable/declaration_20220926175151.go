package main

import "fmt"

var foo string = "Go is awesome"

// 可以只声明、不赋值 -> var foo string
// Go 语言的书写风格是先 变量名 后 变量类型

// 可以同时声明多个变量，多个变量需要使用小括号包裹起来
/* var (foo string = "Hello" 
		bar string ="World"
		baz int = 10086) */

// 如果声明时即赋值,那么可以省略书写数据类型 -> 变量类型将根据传入的值(右边)来自动推断得出
// var foo = "What's my type?"

func main() {
	test := "?"
	// 仅限于函数内,还可以进一步省略 var 关键字 -> 转而使用 := 的语法 (short variable declarations)
    fmt.Println(test)
}
