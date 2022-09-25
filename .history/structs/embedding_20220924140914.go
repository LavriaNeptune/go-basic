package main

func main() {
	type Person struct {
		FirstName, LastName string
		Age                 int
	}
	type SuperHero struct {
		Person
		Alias string
	}
	s := {{"Bruce","Wayne",40},batman}
	fmt.Println(s)
}