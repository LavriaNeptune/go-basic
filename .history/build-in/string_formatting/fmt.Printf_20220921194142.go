package main

import (
	"fmt"
)

func main() {
	name := "golang"

	fmt.Println("What is your name?")
	fmt.Printf("My name is %s", name)
	// fmt.Printf 可以格式化输入内容 -> f 来自 formatting 首字母。
	// 类似 %s 这样通过 % 格式的字符串是所谓的注释动词...用来表示如何对参数进行格式化
}	