package main

import "fmt"

type Car struct {
	Name string
	Year int
}

func (c Car) IsLatest() bool {
	return c.Year >= 2017
}

func main() {
	c := Car{"Tesla", 2022}
	fmt.Println("Return Value:",c.IsLast)
}