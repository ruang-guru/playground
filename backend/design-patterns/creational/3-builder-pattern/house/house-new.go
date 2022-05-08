package house

// Nah untuk builder sendiri kita harus menentukan builder nya itu yang mana. Anggap saja seperti Tukang yang ahli dalam membuat sesuatu
// Disini masih sama kita masih mengambil contoh untuk membuat sebuah Rumah
// Katakan terdapat 2 jenis builder (Tukang). Builder Zimbabwe dan Builder Indonesia

// Kita definisikan dulu interface(kontrak) general untuk builder

type House struct {
	NumOfWindows    int
	NumOfDoors      int
	HasGarage       bool
	HasSwimmingPool bool
}

type HouseBuilder interface {
	buildWindow(numOfWindow int)
	buildDoor()
	buildGarage()
	buildSwimmingPool()
	getHouse() House
}

func NewBuilder(builderType string) HouseBuilder {
	if builderType == "zimbabwe" {
		return &zimbabweHouseBuilder{
			house: House{
				HasSwimmingPool: false,
			},
		} // cek zimbabweHouseBuilder.go
	}
	if builderType == "indonesia" {
		return &indonesiaHouseBuilder{} // cek indonesiaBuilder.go
	}

	return nil
}
