package main

import "errors"

// Menghapus dan mengambil nilai dari stack disebut dengan pop
// Pada operasi pop, kita harus mengecek apakah terdapat elemen pada stack atau tidak.
// Jika tidak ada elemen pada stack maka terdapat error stack underflow.

// Untuk mengecek apakah ada elemen pada stack ada dua pendekatan.
// Pendekatan pertama adalah dengan mengecek apakah panjang dari data adalah 0.
// Pendekatan kedua adalah dengan mengecek apakah nilai top sama dengan -1.

type Stack struct {
	Top  int
	Data []int
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack underflow")
	} else {
		poppedValue := s.Data[s.Top]
		s.Top -= 1
		s.Data = s.Data[:len(s.Data)-1]
		return poppedValue, nil
	}
}
