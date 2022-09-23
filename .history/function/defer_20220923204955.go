package main

import "fmt"


func main() {
	defer fmt.Println("第一个怕丢i")
	fmt.Println(sum)
}