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

type PowerDrawer interface {
	Draw(power int)
}

func (m mobile) Draw(power int) {
	fmt.Printf("%T -> brand:%s,power %d", m, m.brand, power)
}

func (l mobile) Draw(power int) {
	fmt.Printf("%T -> cpu:%s,power %d", l, m.brand, power)
}

func (socket) Plug(device PowerDrawer, power int) {
	device.Draw(power)
}

func main() {
	m := mobile{"Apple"}
	l := laptop{"Intel i9"}

	s := socket{}

	s.Plug(m,10)
	s.Plug(l,50)
}