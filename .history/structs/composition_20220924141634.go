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
	p := Person{"Bruce","Wayne"}

	fmt.Println(s)
}