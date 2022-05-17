package coffe

import (
	"strings"
)

// Nah tetapi ada condiment yang unik disini yaitu Ice.
// Jika sebuah coffe ditambahkan Ice maka akan terdapat Magic didalamnya.
// Magic ini terjadi jika Coffe tersebut berharga diatas 3.30 dan Ice sudah ditambahkan sebelumnya (2x ice)
// Magic tersebut adalah jika Coffe tersebut sudah ditambahkan Ice sebelumnya maka Kopi tersebut menjadi BEKU

type Ice struct {
	Coffe Coffe
}

func (i Ice) GetCost() float64 {
	return i.Coffe.GetCost() + 0.20
}

func (i Ice) GetDescription() string {
	return description // TODO: replace this
}

// check repeating word in a string
// Use this to check repeating word in string ;)
func checkRepeatingWord(s string) bool {
	input := strings.Fields(s)
	for _, word := range input {
		if strings.Count(s, word) > 1 {
			return true
		}
	}
	return false
}
