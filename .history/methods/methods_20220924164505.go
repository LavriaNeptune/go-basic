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

func (c *Car) UpdateName(name string) {
	// 为什么此处要接受的是指针 -> 如果是非指针,c 传入的是元数据的值复制,修改并不会影响到原数据(可以尝试去掉 * 再运行程序)
	c.Name = name
}

func main() {
	c := Car{"Tesla", 2022}
	fmt.Println("Return Value:",c.IsLatest())

	c.UpdateName("BMW")
	// ()
	fmt.Println(c)
}