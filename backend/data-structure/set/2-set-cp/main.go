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

// Delete - menghapus elemen dari set
func (s *Set) Delete(elem string) (bool, error) {
	return fmt.Errorf("replace this with your code") // TODO: replace this
}

// Contains - memeriksa apakah elemen ada dalam set
func (s *Set) Contains(elem string) bool {
	return false // TODO: replace this
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

	_, err := set.Delete("Gine")
	if err != nil {
		fmt.Println(err)
	}

	set.List()

	if set.Contains("Aditira") {
		fmt.Println("Aditira is in the set")
	} else {
		fmt.Println("Aditira is not in the set")
	}
}
