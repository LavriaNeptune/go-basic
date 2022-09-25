package main

func main() {
	type Person struct {
		FirstName string
		LastName  string
		// 同类型字段可以用逗号分隔,简写为:FirstName, LastName string
		Age int
	}
}