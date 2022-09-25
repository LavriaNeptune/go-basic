// 切片是数组的衍生概念 -> 切片提供了更多便利、灵活的功能
package main

import "fmt"

var arr = [6]int{1, 2, 3, 4, 5, 6}

func main() {
	// 完整切片声明方式,因为切片是不定长的
	// var s []T

	// 可以使用 make([]T,len,cap) 来初始化一个切片

	s := arr[1:4]
	// 最常见创建切片方式 -> 通过一个存在的数组创建
	// 常见省略写法 ↓
	// [:3] -> 取数组后两个元素组成切片
	// [2:] -> 取数组后两个元素组成切片
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d", s, len(s), cap(s))
	// length 表示切片的长度 -> len() 可查;cap 表示切片的最大长度 -> cap() 可查
	// 切片是数组底层数组的引用类型数据
}