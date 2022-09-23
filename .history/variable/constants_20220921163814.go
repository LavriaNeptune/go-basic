package main

import "fmt"

var foo string = "Go is awesome"


func main() {
	test := "?"
	// 在函数内,还可以进一步省略 var 关键字 -> 转而使用 := 的语法
    fmt.Println(test)
}