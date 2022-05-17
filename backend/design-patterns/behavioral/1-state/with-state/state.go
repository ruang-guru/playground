package withstate

// Pada state pattern kita akan mendefine sebuah "State"
type State interface {
	Press()
	CanTurnOnLaptop() bool
}

func NewLaptop() *Laptop {
	l := &Laptop{}
	// State awal laptop adalah Off
	l.ChangeState(&Off{l})
	return l
}

// Selanjutkan kita akan taruh State itu didalam sebuah laptop ini yang kita sebut "Composition"
type Laptop struct {
	State State
}

// Ketika kita memanggil Press(), Press() kita delegasikan ke State nya untuk mengaturnnya
func (l Laptop) Press() {
	l.State.Press()
}

// Samah hal nya juga untuk CanTurnOnLaptop kita delegasikan call nya ke State nya
func (l Laptop) CanTurnOnLaptop() bool {
	return l.State.CanTurnOnLaptop()
}

func (l *Laptop) ChangeState(state State) {
	l.State = state
}

// nah bayangkan disini kita juga mempunyai func untuk ngecek baterai yang dipunya sama laptop ny
// Ini kita akan gunakan untuk melakukan logic pada State yang kita buat
func (l Laptop) IsThereBattery() bool {
	return true
}

// Ini adalah concrete implementation dari state On
type On struct {
	Laptop *Laptop
}

func (o On) Press() {
	o.Laptop.ChangeState(&Off{o.Laptop})
}

func (o On) CanTurnOnLaptop() bool {
	return true
}

// Ini adalah concrete implementation dari state Off
type Off struct {
	Laptop *Laptop
}

func (o Off) Press() {
	if !o.Laptop.IsThereBattery() {
		return
	}
	o.Laptop.ChangeState(&On{o.Laptop})
}

func (o Off) CanTurnOnLaptop() bool {
	return false
}

// Nah kan sudah selesai tuh, terus teman teman akan mencoba sendiri menambahkan state baru yaitu Sleep. Cek 1-state-cp yaaa
