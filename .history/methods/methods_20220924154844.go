package main

import "fmt"

type Car struct {
	Name string
	Year int
}

func (c Car) IsLatest() bool {
	return c.Year >= 2017
}
// 本质上方法就是一个带有特殊接收参数的函数...

func (c Car) UpdateName(name string) {
	c.Name = name
}

func main() {
	c := Car{"Tesla", 2022}
	fmt.Println("Return Value:",c.IsLatest())
}