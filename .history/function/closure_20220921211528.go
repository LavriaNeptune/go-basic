package main

func myFunction() {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}