package main

import "fmt"

type mobile struct {
	brand string
}

type laptop struct {
	cpu string
}

type toaster struct {
	amount int
}

type kettle struct {
	quantity string
}

type socket struct{}

func (m mobile) Draw(power int) {
	fmt.Printf("%T -> brand:%s,power %d", m, m.brand, power)
}
func (socket) Plug(device mobile, power int) {
	device.Draw(power)
}

func main() {
	m := mobile{"Apple"}
	l :=laptop{"Intel i9"}

	s := socket{}

	s.Plug(m)
}