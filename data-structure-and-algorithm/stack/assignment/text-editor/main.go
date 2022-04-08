package main

import (
	"fmt"

	"github.com/ruang-guru/playground/data-structure-and-algorithm/stack/assignment/text-editor/stack"
)

// Implementasikan teks editor sederhana
// Teks editor ini memiliki 4 method
// Write: digunakan untuk menulis per karakter
// Read: digunakan untuk mencetak karakter yang telah ditulis
// Undo: digunakan untuk melakukan operasi undo
// Redo: digunakan untuk melakukan operasi redo

type TextEditor struct {
	UndoStack *stack.Stack
	RedoStack *stack.Stack
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		UndoStack: stack.NewStack(),
		RedoStack: stack.NewStack(),
	}
}

func (te *TextEditor) Write(ch rune) {
	//beginanswer
	te.RedoStack.SetToEmpty()
	te.UndoStack.Push(ch)
	//endanswer
}

func (te *TextEditor) Read() []rune {
	//beginanswer
	tempStack := stack.NewStack()
	var output []rune
	for te.UndoStack.Top > -1 {
		ch, _ := te.UndoStack.Pop()
		tempStack.Push(ch)
	}
	for tempStack.Top > -1 {
		ch, _ := tempStack.Peek()
		fmt.Printf("%s", string(ch))
		output = append(output, ch)
		te.UndoStack.Push(ch)
		tempStack.Pop()
	}

	return output
	//endanswer
}

func (te *TextEditor) Undo() {
	//beginanswer
	ch, err := te.UndoStack.Peek()
	if err != nil {
		return
	}
	te.UndoStack.Pop()
	te.RedoStack.Push(ch)
	//endanswer
}

func (te *TextEditor) Redo() {
	//beginanswer
	ch, err := te.RedoStack.Peek()
	if err != nil {
		return
	}
	te.RedoStack.Pop()
	te.UndoStack.Push(ch)
	//endanswer
}
