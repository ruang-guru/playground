package main

import (
	"errors"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

var ErrStackOverflow = errors.New("stack overflow")

type Stack struct {
	//beginanswer
	Top  int
	Size int
	Data []int
	//endanswer
}

func NewStack(size int) Stack {
	//beginanswer
	return Stack{
		Top:  -1,
		Size: size,
		Data: []int{},
	}
	//endanswer
}

func (s *Stack) Push(Elemen int) error {
	//beginanswer
	if len(s.Data) == s.Size {
		return ErrStackOverflow
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
	}
	return nil
	//endanswer
}
