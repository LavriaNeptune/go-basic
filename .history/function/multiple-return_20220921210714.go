package main

import "fmt"

func Test(str string)(string,int) {
	return str,996
	// Go 语言的 return 可以返回多个值,值与值之间通过逗号间隔开即可...
}

// 通过 func 关键字以如上格式声明函数 -> 所声明的函数可以在其他函数中进行调用...

func main() {
	msg,num := Test("Hello")
	fmt.Println(msg,num)
}