package main

import (
	"fmt"
)

func main() {
	type MyAliasSting = string
	// 使用 type 关键字,将 MyAliasString
	var str MyAliasSting ="hello Go"

	fmt.Printf(str)
}