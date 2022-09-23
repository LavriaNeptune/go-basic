package main

import "fmt"

func Add(str string){
	fmt.Println(str)
}
// 通过 func 关键字以如上格式声明函数 -> 所声明的函数可以在其他函数中进行调用...

func main() {
	myFunction("Hello")
}