package main

func main() {
	type Person struct {
		FirstName string
		LastName  string
		// 同类型字段可以简写为:
		Age       int
	}
}