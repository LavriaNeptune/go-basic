package main

import "fmt"


func main() {
	defer fmt.Println("第一个排队")
	defer fmt.Println("第二个排队")
	// defer 可以延迟执行同一层级的函数...并遵循 后入先出 的执行模式
	fmt.Println("Doing some work")
}