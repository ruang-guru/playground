package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi peek.

var ErrEmptyStack = errors.New("stack is empty")

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

func (s *Stack) Peek() (int, error) {
	//beginanswer
	if s.Top == -1 {
		return 0, ErrEmptyStack
	} else {
		return s.Data[s.Top], nil
	}
	//endanswer
}
