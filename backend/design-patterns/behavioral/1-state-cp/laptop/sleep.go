package laptop

type Sleeping struct {
	Laptop *Laptop
}

func (s Sleeping) Press() {
	// TODO: answer here
}

func (s Sleeping) CanTurnOnLaptop() bool {
	return true
}

func (s Sleeping) Sleep() {
	// TODO: answer here
}
