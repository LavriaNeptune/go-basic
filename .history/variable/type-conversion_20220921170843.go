package main

import "fmt"



func main() {
	i := 42
	f := float64(i)
	u := uint(f)
	// 通过与类型名称一致的内建类型转换函数进行类型转换
	fmt.Printf("%T %T", f, u)
}