package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi peek.

var ErrEmptyStack = errors.New("stack is empty")

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

func (s *Stack) Peek() (int, error) {
	// TODO: answer here
	if s.Top == -1 {
		return 0, ErrEmptyStack
	}

	return s.Data[s.Top], nil
}
