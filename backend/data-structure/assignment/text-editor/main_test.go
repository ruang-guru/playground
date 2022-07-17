package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Text Editor", func() {
	Describe("Write", func() {
		Describe("Write a character", func() {
			It("Should append a character", func() {
				te := NewTextEditor()
				te.Write('a')
				Expect(te.UndoStack.Data).To(HaveLen(1))
				Expect(te.UndoStack.Data[0]).To(Equal('a'))

				te.Write('b')
				te.Write('c')
				Expect(te.UndoStack.Data).To(HaveLen(3))
				Expect(te.UndoStack.Data[0]).To(Equal('a'))
				Expect(te.UndoStack.Data[1]).To(Equal('b'))
				Expect(te.UndoStack.Data[2]).To(Equal('c'))
			})
		})
		Describe("Write multiple characters", func() {
			It("Should append multiple characters", func() {
				te := NewTextEditor()
				te.Write('h')
				te.Write('e')
				te.Write('y')
				Expect(te.UndoStack.Data).To(HaveLen(3))
				Expect(te.UndoStack.Data[0]).To(Equal('h'))
				Expect(te.UndoStack.Data[1]).To(Equal('e'))
				Expect(te.UndoStack.Data[2]).To(Equal('y'))
			})
		})

	})
	Describe("Read", func() {
		Describe("Read on text editor with inserted characters", func() {
			It("Should print inserted characters", func() {
				te := NewTextEditor()
				te.Write('h')
				te.Write('e')
				te.Write('l')
				te.Write('l')
				te.Write('o')

				output := te.Read()
				Expect(output).To(Equal([]rune{'h', 'e', 'l', 'l', 'o'}))
			})
		})
		Describe("Read on text editor with no inserted characters", func() {
			It("Should print empty", func() {
				te := NewTextEditor()

				output := te.Read()
				Expect(output).To(BeNil())
			})
		})

	})
	Describe("Undo", func() {
		Describe("Undo on empty state text editor", func() {
			It("Undo nothing", func() {
				te := NewTextEditor()
				te.Undo()
				Expect(te.UndoStack.IsEmpty()).To(BeTrue())
				Expect(te.RedoStack.IsEmpty()).To(BeTrue())
			})
		})
		Describe("Undo on write first character", func() {
			It("Should undo the writing and push element to redo stack", func() {
				te := NewTextEditor()
				te.Write('a')
				te.Undo()
				Expect(te.UndoStack.IsEmpty()).To(BeTrue())
				Expect(te.RedoStack.IsEmpty()).To(BeFalse())
				Expect(te.RedoStack.Data).To(HaveLen(1))
			})
		})
		Describe("Undo on write multiple characters", func() {
			It("Should undo the writing and push element to redo stack", func() {
				te := NewTextEditor()
				te.Write('w')
				te.Write('o')
				te.Write('r')
				te.Write('l')
				te.Write('d')
				te.Undo()
				Expect(te.UndoStack.IsEmpty()).To(BeFalse())
				Expect(te.RedoStack.IsEmpty()).To(BeFalse())
				Expect(te.UndoStack.Data).To(HaveLen(4))
				Expect(te.RedoStack.Data).To(HaveLen(1))
			})
		})
		Describe("Multiple undo on write multiple characters", func() {
			It("Should undo the writing and push element to redo stack", func() {
				te := NewTextEditor()
				te.Write('w')
				te.Write('o')
				te.Write('r')
				te.Write('l')
				te.Write('d')
				te.Undo()
				te.Undo()
				Expect(te.UndoStack.IsEmpty()).To(BeFalse())
				Expect(te.RedoStack.IsEmpty()).To(BeFalse())
				Expect(te.UndoStack.Data).To(HaveLen(3))
				Expect(te.RedoStack.Data).To(HaveLen(2))
			})
		})
	})
	Describe("Redo", func() {
		Describe("Redo on empty state text editor", func() {
			It("Redo nothing", func() {
				te := NewTextEditor()
				te.Redo()
				Expect(te.UndoStack.IsEmpty()).To(BeTrue())
				Expect(te.RedoStack.IsEmpty()).To(BeTrue())
			})
		})
		Describe("Redo after undo on write first character", func() {
			It("Should redo the writing", func() {
				te := NewTextEditor()
				te.Write('a')
				te.Undo()
				te.Redo()
				Expect(te.UndoStack.IsEmpty()).To(BeFalse())
				Expect(te.RedoStack.IsEmpty()).To(BeTrue())
				Expect(te.UndoStack.Data).To(HaveLen(1))
			})
		})
		Describe("Redo after undo on write multiple characters", func() {
			It("Should redo the writing", func() {
				te := NewTextEditor()
				te.Write('w')
				te.Write('o')
				te.Write('r')
				te.Write('l')
				te.Write('d')
				te.Undo()
				te.Redo()
				Expect(te.UndoStack.IsEmpty()).To(BeFalse())
				Expect(te.RedoStack.IsEmpty()).To(BeTrue())
				Expect(te.UndoStack.Data).To(HaveLen(5))
			})
		})
		Describe("Multiple redo after multiple undo on write multiple characters", func() {
			It("Should redo the writing", func() {
				te := NewTextEditor()
				te.Write('w')
				te.Write('o')
				te.Write('r')
				te.Write('l')
				te.Write('d')
				te.Undo()
				te.Undo()
				te.Undo()
				te.Redo()
				te.Redo()
				Expect(te.UndoStack.IsEmpty()).To(BeFalse())
				Expect(te.RedoStack.IsEmpty()).To(BeFalse())
				Expect(te.UndoStack.Data).To(HaveLen(4))
				Expect(te.RedoStack.Data).To(HaveLen(1))
			})
		})
	})
})
