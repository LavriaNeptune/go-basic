package main

import "fmt"

func Test(str string)(string,int) {
	return str,996
	// 可以使用 return 语句为函数返回一个结果值 -> 返回值时需要给返回值设定返回的数据类型
}

// 通过 func 关键字以如上格式声明函数 -> 所声明的函数可以在其他函数中进行调用...

func main() {
	msg,num := Test("Hello")
	fmt.Println(msg)
}