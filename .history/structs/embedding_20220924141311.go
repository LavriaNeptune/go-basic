package main

import "fmt"

func main() {
	type Person struct {
		FirstName, LastName string
		Age                 int
	}
	type SuperHero struct {
		Person
		Alias string
	}
	s := {}

	s.FirstName = "Bruce"
	s.LastName = "Wayne"
	s.Age = 40
	s.Alias = "batman"

	fmt.Println(s)
}