package main

import "fmt"

func main() {
	day := "monday"
	switch day{
		// 省略 switch 条件 -> 表示条件始终为 true
		case "monday":
			fmt.Println("Time to work!")
		case "Saturday":
			fmt.Println("Third")
		case "friday":
			fmt.Println("let's party")
		
		default:
			fmt.Println("other case")
	}
	// Go 语言中 switch 语句和 if 语句类似,都是可以省略掉条件的小括号的...可以用来替代 if 嵌套层级过多的情况
}