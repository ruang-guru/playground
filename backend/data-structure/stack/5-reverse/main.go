package main

import (
	"fmt"
)

type Stack struct {
	Top  int
	Data []int
}

func (s *Stack) Push(Elemen int) {
	s.Top += 1
	s.Data = append(s.Data, Elemen)
}

func (s *Stack) Pop() (int, error) {
	if s.Top == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	s.Top -= 1
	return s.Data[s.Top], nil
}

func NewStack() *Stack {
	return &Stack{
		Top:  0,
		Data: []int{},
	}
}

func main() {
	stack := NewStack()
	word := "hello world"
	for _, char := range word {
		stack.Push(int(char))
	}
	var reversed string
	for stack.Top > 0 {
		char, _ := stack.Pop()
		reversed += string(rune(char))
	}
	fmt.Println(reversed)
}
