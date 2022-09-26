package main

import "fmt"

const foo string = "Hello"

// 通过 const 定义的变量后续不可更改变量的值 -> 所以又称常量...PS:常量不可通过 := 

func main() {
	// ↓ 尝试修改常量 foo 的值将会报错:cannot assign to foo
	// foo = 1;

  fmt.Println(foo)
}