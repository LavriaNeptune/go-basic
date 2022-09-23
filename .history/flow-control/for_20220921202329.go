// Go 语言的循环语句只有 for 一种(很聪明,选取了最好用的一种)

package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// 类似于其他语言中 for 循环语句...
}