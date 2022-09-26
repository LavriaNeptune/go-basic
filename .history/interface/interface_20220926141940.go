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
	fmt.Printf("%T -> brand:%s,")
	
}