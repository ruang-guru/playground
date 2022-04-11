package main

import "errors"

// Melihat elemen teratas pada stack disebut dengan peek.

type Stack struct {
	Top  int
	Data []int
}

func (s *Stack) Peek() (int, error) {
	if s.Top == -1 {
		return 0, errors.New("stack is empty")
	} else {
		return s.Data[s.Top], nil
	}
}
