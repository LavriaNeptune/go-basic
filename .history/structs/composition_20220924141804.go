package main

import "fmt"

func main() {
	type Person struct {
		FirstName, LastName string
		Age                 int
	}
	type SuperHero struct {
		Person Person
		Alias string
	}
	// 组合方式,比单纯嵌入更易读...直接将组合类型定义为嵌套字段
	p := Person{"Bruce","Wayne",40}
	s := SuperHero{p,"batman"}

	fmt.Println(s)
}