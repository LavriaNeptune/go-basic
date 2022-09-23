package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x > 5")
	} else if x > 10 {
		fmt.Println("x > 10")
	}else {
		fmt.Println("else case")
	}
	// 可以看出 Go 语言中 if 语句是可以省略掉条件
}