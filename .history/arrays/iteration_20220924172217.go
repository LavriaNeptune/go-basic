package main

import "fmt"

func main(){
	var arr [5]int
	// 声明和其他类型的变量大同小异,[T] 表示类型为数组,T -> 表示数组长度,后接类型表示数组元素的类型
	fmt.Println(arr) // [0 0 0 0 0]

	arr = [5]int{1,2,3,4,5}
	// Go 的数组表示有点诡异,不是 [1,2,3,4,5] 而是  [5]int{1,2,3,4,5} 表示数组、长度5、类型为 int、数组元素值由 {} 包裹着数据依次给出
	fmt.Println(arr) // [1 2 3 4 5]
	
	fmt.Println(arr[2]) // 3
}