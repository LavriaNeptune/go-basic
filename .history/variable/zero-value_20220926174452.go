// 仅声明,未定义的值将会被自动赋予类型相关的 零值 -> 数字类型的零值是 0;字符串类型的零值是空字符串 "";布尔类型的零值是 false

package main

import "fmt"

var ( i int
			f float64
			b bool
			s string)

func main() {
    fmt.Printf("%v %v %v %q\n", i, f, b, s) // 0 0 false ""
}