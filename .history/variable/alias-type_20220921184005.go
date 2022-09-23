package main

import (
	"fmt"
)

func main() {
	type MyAliasSting = string
	var str MyAliasSting ="hello Go"
	// jiang
	fmt.Printf(str)
}