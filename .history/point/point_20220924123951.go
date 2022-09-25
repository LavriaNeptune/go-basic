package main

import "fmt"

// 

func main() {
	var p *int
	// 在基础类型前添加 * -> 获取基本类型所对应的指针类型
	fmt.Println(p) // <nil> -> 指针类型的零值
}