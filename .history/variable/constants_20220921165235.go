package main

import "fmt"

const foo string = "Hello"

// PS:函数外部不可以修改变量值

func main() {
	// ↓ 尝试修改常量 foo 的值将会报错:cannot assign to foo
	// foo = 1;

    fmt.Println(foo)
}