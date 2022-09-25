package main

import "fmt"

// 指针就是一种专门用来存储其他变量的地址的变量

func main() {
	var p *int
	// 在基础类型前添加 * -> 获取基本类型所对应的指针类型
	fmt.Println(p) // <nil> -> 指针类型的零值(Go 的预声明标识符...还可以用来表示接口、通道、切片..的零值)

	a:=10086
	p = &a
	// 通过 
	fmt.Println(p)
}