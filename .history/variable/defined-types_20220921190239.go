package main

import (
	"fmt"
)

func main() {
	type MyDefString string
	// 使用 type 关键字,将 MyDefString 设置为关键字的 string 的同义词
	var str MyAliasString ="hello Go"

	fmt.Printf(str)
}