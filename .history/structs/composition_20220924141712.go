package main

import "fmt"

func main() {
	type Person struct {
		FirstName, LastName string
		Age                 int
	}
	type SuperHero struct {
		Person Person
		Alias string
	}
	// 组合方式
	p := Person{"Bruce","Wayne",40}
	s := SuperHero{p,"batman"}

	fmt.Println(s)
}