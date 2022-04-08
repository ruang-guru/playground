package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi pop.

var ErrStackUnderflow = errors.New("stack underflow")

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
		return errors.New("stack overflow")
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
	}
	return nil
	//endanswer
}

func (s *Stack) Pop() (int, error) {
	//beginanswer
	if s.Top == -1 {
		return 0, ErrStackUnderflow
	} else {
		poppedValue := s.Data[s.Top]
		s.Top -= 1
		s.Data = s.Data[:len(s.Data)-1]
		return poppedValue, nil
	}
	//endanswer
}
