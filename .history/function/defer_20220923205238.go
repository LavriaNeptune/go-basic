package main

import "fmt"


func main() {
	defer fmt.Println("第一个排队")
	fmt.Println(sum)
}