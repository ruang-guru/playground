package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi pop.

var ErrStackUnderflow = errors.New("stack underflow")

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
		return errors.New("stack overflow")
	}

	s.Top++
	s.Data = append(s.Data, Elemen)
	return nil
}

func (s *Stack) Pop() (int, error) {
	// TODO: answer here

	if s.Top < 0 {
		return 0, ErrStackUnderflow
	}

	res := s.Data[s.Top]
	s.Top--
	s.Data = s.Data[0 : len(s.Data)-1]

	return res, nil
}
