// Go 语言的映射的 k 和 v 可以是任何类型
package main

import "fmt"

var m =  map[string]int{
	"a":1,
	"b":2,
}
// 以上是通过 map 字面量初始化一个 map
// 还可以通过 make(map[string]int) 创建对应的空的 map 数据...


func main() {
	fmt.Println(m) // map[a:1 b:2]

	m["c"] = 3
	// 向 map 结构添加数据,同时也用来更新 map 的数据
	fmt.Println(m) 

	fmt.Println("The value of key \"c\":",m["c"])

	d,ok := m["d"]
	fmt.Println("The value of key \"d\":",d,ok)
	// 取值时,map 还会返回一个表示所取值是否存在的布尔值

	for key,value := range m {
		fmt.Println("Key:%")
	}
}