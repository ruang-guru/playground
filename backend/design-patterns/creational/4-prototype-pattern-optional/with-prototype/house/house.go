package house

// Nah untuk menggunakan protpye sendiri ini biasanya kita menggunakan interface yang mereturn dirinya sendiri
// Seperti ini

type Prototype interface {
	Clone() Prototype
}

type House struct {
	NumOfWindows int
	NumOfDoors   int
	typeOfJamban string
	typeOfLamp   string
}

// Nah function NewHouse ini akan dipanggil oleh client yang ingin meng-copy object House
// Perlu diperhatikan NewHouse menerima interace Prototype yang berisi method Clone
func CloneHouse(p Prototype) Prototype {
	return p.Clone()
}

func NewHouse(numWindows, numDoors int, typeOfJamban, typeOfLamp string) *House {
	return &House{
		NumOfWindows: numWindows,
		NumOfDoors:   numDoors,
		typeOfJamban: typeOfJamban,
		typeOfLamp:   typeOfLamp,
	}
}

func (h *House) Clone() Prototype {
	return &House{
		NumOfWindows: h.NumOfWindows,
		NumOfDoors:   h.NumOfDoors,
		typeOfJamban: h.typeOfJamban,
		typeOfLamp:   h.typeOfLamp,
	}
}
