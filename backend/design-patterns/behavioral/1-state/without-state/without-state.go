package withoutstate

// Ini adalah contoh sederhana ketika kita tidak menggunakan state pattern

// Bayangin kita punya Laptop/PC

type LaptopState string

const (
	On  LaptopState = "on"
	Off LaptopState = "off"
	Dim LaptopState = "dim"
)

// Dimana kita mempunyai property State yang menandakan State On/Off
type Laptop struct {
	State LaptopState
}

func (l Laptop) Press() {
	if l.State == Off {
		l.State = On
	} else if l.State == On {
		l.State = Off
	} else {
		l.State = Dim
	}
}

// Anggaplah kitap unya satu func yang nanyain kalau Laptop nya bisa nyala
func (l Laptop) CanTurnOnLaptop() bool {
	return l.State == On
}

// Tapi apa yang terjadi ketika kita menambahkan state pada enum diatas ? menjadi "Sleep". kita pasti akan menambahkan if-else pada function press kan
// Nah dengan state pattern kode kita dapat menjadi lebih bersih dan mengikuti prinsip Solid
