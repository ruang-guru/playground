package main

import (
	"fmt"
)

// Set - mewakili set dari data structure yang ditetapkan
type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	return &set
}

// Add - menambahkan elemen ke dalam set
func (s *Set) Add(elem string) {
	s.Elements[elem] = struct{}{}
}

// List - menampilkan elemen dari set
func (s *Set) List() {
	for k := range s.Elements {
		fmt.Println(k)
	}
}

func main() {
	fmt.Println("Sets Demo -> List of Set:")
	set := NewSet()

	set.Add("Aditira")
	set.Add("Gina")
	set.Add("Dito")
	set.Add("Aditira")
	set.Add("Imron")
	set.Add("Imron")
	
	set.List()
}
