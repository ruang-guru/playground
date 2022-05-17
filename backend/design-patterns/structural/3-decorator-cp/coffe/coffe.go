package coffe

// Katakanlah kita akan membuat sebuah object Kopi
// Coffe tersebut bisa kita dekorasi dengan berbagai macam Condiment, seperti Whipcream, Mocha, dll

// Kita mulai dengan interface yang akan kita penuhi

type Coffe interface {
	GetCost() float64
	GetDescription() string
}

type Espresso struct {
}

func (e Espresso) GetCost() float64 {
	return 2.00
}

func (e Espresso) GetDescription() string {
	return "Espresso"
}

type Coldbrew struct {
}

func (c Coldbrew) GetCost() float64 {
	return 3.00
}

func (c Coldbrew) GetDescription() string {
	return "Coldbrew"
}
