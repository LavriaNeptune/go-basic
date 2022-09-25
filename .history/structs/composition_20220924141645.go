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
	p := Person{"Bruce","Wayne",40}
	s:=SuperHero

	fmt.Println(s)
}