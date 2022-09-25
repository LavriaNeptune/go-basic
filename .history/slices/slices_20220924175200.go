// 切片是数组的衍生概念 -> 切片提供了更多便利、灵活的功能
package main

import "fmt"

var arr = [6]int{1, 2, 3, 4, 5, 6}

func main() {
	// 完整切片声明方式,因为切片是不定长的
	// var s []T

	// 可以使用 make([]T,)

	s := arr[1:4]
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d", s, len(s), cap(s))
	// length 表示切片的长度 -> len() 可查;cap 表示切片的最大长度 -> cap() 可查
	// 切片是数组底层数组的引用类型数据
}