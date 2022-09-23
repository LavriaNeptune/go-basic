// Go 语言中,仅声明,但未定义的值将会被赋予类型相关的 零值 -> 数字类型的零值是 0;字符串类型的零值是空字符串 "";布尔类型的零值是 false

package main

import "fmt"

const foo string = "Hello"

// PS:函数外部不可以修改变量的值...

func main() {
	// ↓ 尝试修改常量 foo 的值将会报错:cannot assign to foo
	// foo = 1;

    fmt.Println(foo)
}