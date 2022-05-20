package laptop

// Pada state pattern kita akan mendefine sebuah "State"
type State interface {
	Press()
	Sleep()
	CanTurnOnLaptop() bool
}

func New() *Laptop {
	l := &Laptop{}
	// State awal laptop adalah Off
	l.CurrentState = "Off"
	l.ChangeState(Off{l})
	return l
}

// Selanjutkan kita akan taruh State itu didalam sebuah laptop ini yang kita sebut "Composition"
type Laptop struct {
	State        State
	CurrentState string // For easier debugging and writing testing
}

// Ketika kita memanggil Press(), Press() kita delegasikan ke State nya untuk mengaturnnya
func (l Laptop) Press() {
	l.State.Press()
}

// Sama hal nya juga untuk CanTurnOnLaptop kita delegasikan call nya ke State nya
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

func (l Laptop) GetState() State {
	return l.State
}

func (l Laptop) Sleep() {
	l.State.Sleep()
}

func (l *Laptop) ChangeCurrentState(state string) {
	l.CurrentState = state
}
