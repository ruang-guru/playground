package laptop

// Ini adalah concrete implementation dari state Off
type Off struct {
	Laptop *Laptop
}

func (o Off) Press() {
	if !o.Laptop.IsThereBattery() {
		return
	}
	// chane laptop currentstate to "On" by memory
	o.Laptop.ChangeState(&On{o.Laptop})
}

func (o Off) CanTurnOnLaptop() bool {
	return false
}

func (o Off) Sleep() {
	return
}
