package main

import "fmt"

func main() {
	var p *int
	// 在基础类型前添加 * -> 获取对应基本类型的
	fmt.Println(p) // <nil> -> 指针类型的零值
}