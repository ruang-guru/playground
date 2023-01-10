package laptop

type Sleeping struct {
	Laptop *Laptop
}

func (s Sleeping) Press() {
	// TODO: answer here
	if !s.Laptop.IsThereBattery() {
		return
	}
	s.Laptop.ChangeState(&On{s.Laptop})
}

func (s Sleeping) CanTurnOnLaptop() bool {
	return true
}

func (s Sleeping) Sleep() {
	// TODO: answer here
	return
}
