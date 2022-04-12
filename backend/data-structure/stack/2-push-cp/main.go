package main

import (
	"errors"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

var ErrStackOverflow = errors.New("stack overflow")

type Stack struct {
	// TODO: answer here
	Size int
	Top  int
	Data []int
}

func NewStack(size int) Stack {
	// TODO: answer here

	s := Stack{}
	s.Data = make([]int, 0)
	s.Size = size
	s.Top = -1

	return s
}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
	if s.Top+1 >= s.Size {
		return ErrStackOverflow
	}

	s.Top++
	s.Data = append(s.Data, Elemen)
	return nil
}
