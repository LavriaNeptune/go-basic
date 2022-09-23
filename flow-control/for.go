// Go 语言的循环语句只有 for 一种(很聪明,选取了最好用的一种)

package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		// 省略掉 for 循环的语句表示无限循环
		fmt.Println(i)
	}
	// 类似于其他语言中 for 循环语句...不过 Go 语言中的 for 语句也是可以将条件的小括号省略掉的
}