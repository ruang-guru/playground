package virtual_house

import "github.com/ruang-guru/playground/backend/design-patterns/creational/4-prototype-pattern-optional/without-prototype/house"

type VirtualHouse struct {
}

func (f VirtualHouse) CopyHouse(h house.House) *house.House {
	// bisa dilihat disini kita tidak bisa melakukan copy secara sepenuhnya terhadap object rumah
	// Karena ada beberapa field yang private.
	return &house.House{
		NumOfWindows: h.NumOfWindows,
		NumOfDoors:   h.NumOfDoors,
	}
}
