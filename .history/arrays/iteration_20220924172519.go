package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index:%d,Element:%d\n",i,arr[i])
	}
	// 直接使用 for 循环迭代出数组的每个值...len() 函数用来获取数组长度

	for i,e := range arr{
		
	}
}