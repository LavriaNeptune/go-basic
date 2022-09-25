package main

import "fmt"

func main() {
	type Person struct {
		FirstName string
		LastName  string
		// 同类型字段可以用逗号分隔,简写为:FirstName, LastName string
		Age int
	}
	// 以上代码定义一个名为 Person 的结构类型的变量

	var p1 Person
	fmt.Println("Person 1",p1)
	// 输入为 Person 1 {  0} -> 结构类型字段值都是以对应类型的零值进行初始化的:所以结果就是 "" "" 0

	p2 := Person{FirstName:"L"}
}