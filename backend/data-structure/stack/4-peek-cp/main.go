package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi peek.

var ErrEmptyStack = errors.New("stack is empty")

type Stack struct {
	// TODO: answer here
}

func NewStack(size int) Stack {
	// TODO: answer here
}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
}

func (s *Stack) Peek() (int, error) {
	// TODO: answer here
}
