package main

import "fmt"

func main(){
	var arr [5]int
	// 声明和其他类型的变量大同小异,[T] 表示类型为数组,T -> 表示数组长度,后接类型表示数组元素的类型
	fmt.Println(arr) // [0 0 0 0 0]

	arr = [5]int{1,2,3,4,5}

}