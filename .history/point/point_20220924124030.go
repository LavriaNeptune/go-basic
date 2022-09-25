package main

import "fmt"

// 指针就是一种专门用来存储其他变量的地址的便来纪念馆

func main() {
	var p *int
	// 在基础类型前添加 * -> 获取基本类型所对应的指针类型
	fmt.Println(p) // <nil> -> 指针类型的零值
}