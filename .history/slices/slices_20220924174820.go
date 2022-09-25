// 切片是数组的衍生概念 -> 切片提供了更多便利、灵活的功能
package main

import "fmt"

var arr = [5]int{1, 2, 3, 4, 5,6}

func main() {
	s := arr[1:4]
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d", s, len(s), cap(s))

}