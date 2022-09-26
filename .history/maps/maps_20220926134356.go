// Go 语言的映射的 k 和 v 可以是任何类型
package main

import "fmt"

var m =  map[string]int{
	"a":1,
	"b":2,
}
// 还可以通过 make(map[string]int) 创建对应的空的 map 数据...


func main() {
	fmt.Println(m) // map[a:1 b:2]
}