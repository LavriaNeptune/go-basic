package main

import (
	"fmt"
)

func main() {
	name := "golang"

	fmt.Println("What is your name?")
	fmt.Printf("My name is %s", name)
	// fmt.Printf 函数在 fmt.Println 函数的基础上添加了上了换行符...并且在每个参数之间都添加了空格作为间隔
}