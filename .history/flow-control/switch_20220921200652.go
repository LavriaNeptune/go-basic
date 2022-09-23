package main

import "fmt"

func main() {
	day := "monday"
	switch day{
		case "monday":
			fmt.Println("Time to work!")
		case "friday":
			fmt.Println("let's party")
		default:
			fmt.Println("other case")
	}
	
	// 可以看出 Go 语言中 if 语句是可以省略掉条件的小括号的...
}