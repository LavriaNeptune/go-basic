package main

import (
	"fmt"
)

func main() {
	type MyAliasSting = string
	var str MyAliasSting ="hello Go"
	// 将
	fmt.Printf(str)
}