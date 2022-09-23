package main

import (
	"fmt"
)

func main() {
	type MyDefString string
	// 使用 type 关键字,将 MyDefString 类型通过 string 进行定义
	var str MyDefString ="Defined Types"

	fmt.Printf("%T - %s", str, str) //main.MyDefString - Defined Types
}