const { expect, describe, it } = require('@jest/globals');
const TextEditor = require('./index')

describe("TextEditor", () => {
    describe("Write", () => {
        describe("Write a character", () => {
            it("should append a character", () => {
                let te = new TextEditor()
                te.write("a")
                expect(te.undoStack.length).toEqual(1)
                expect(te.undoStack[0]).toEqual("a")

                te.write("b")
                te.write("c")
                expect(te.undoStack.length).toEqual(3)
                expect(te.undoStack[0]).toEqual("a")
                expect(te.undoStack[1]).toEqual("b")
                expect(te.undoStack[2]).toEqual("c")
            })
        })

        describe("Write multiple characters", () => {
            it("should append multiple characters", () => {
                let te = new TextEditor()
                te.write("h")
                te.write("e")
                te.write("y")
                expect(te.undoStack.length).toEqual(3)
                expect(te.undoStack[0]).toEqual("h")
                expect(te.undoStack[1]).toEqual("e")
                expect(te.undoStack[2]).toEqual("y")
            })
        })
    })
    describe("Read", () => {
        describe("Read on text editor with inserted characters", () => {
            it("should return inserted characters", () => {
                let te = new TextEditor()
                te.write("h")
                te.write("e")
                te.write("l")
                te.write("l")
                te.write("o")
                
                let output = te.read()
                expect(output).toEqual("hello")
            })
        })

        describe("Read on text editor with no inserted characters", () => {
            it("should return empty", () => {
                let te = new TextEditor()
                let output = te.read()
                expect(output).toEqual("")
            })
        })
    })
    describe("Undo", () => {
        describe("Undo on empty state text editor", () => {
            it("should undo nothing", () => {
                let te = new TextEditor()
                te.undo()

                expect(te.undoStack.length).toEqual(0)
                expect(te.redoStack.length).toEqual(0)
                
            })
        })
        describe("Undo on write first character", () => {
            it("should undo the writing and push element to redo stack", () => {
                let te = new TextEditor()
                te.write("a")
                te.undo()
                
                expect(te.undoStack.length).toEqual(0)
                expect(te.redoStack.length).toEqual(1)
            })
        })
        describe("Undo on write multiple characters", () => {
            it("should undo the writing and push element to redo stack", () => {
                let te = new TextEditor()
                te.write('w')
				te.write('o')
				te.write('r')
				te.write('l')
				te.write('d')
                te.undo()
                let output = te.read()

                expect(output).toEqual("worl")
                expect(te.undoStack.length).toEqual(4)
                expect(te.redoStack.length).toEqual(1)
            })
        })
        describe("Multiple undo on write multiple characters", () => {
            it("should undo the writing and push element to redo stack", () => {
                let te = new TextEditor()
                te.write('w')
				te.write('o')
				te.write('r')
				te.write('l')
				te.write('d')
                te.undo()
                te.undo()

                let output = te.read()

                expect(output).toEqual("wor")

                expect(te.undoStack.length).toEqual(3)
                expect(te.redoStack.length).toEqual(2)
            })
        })
    })
    
    describe("Redo", () => {
        describe("Redo on empty state text editor", () => {
            it("should redo nothing", () => {
                let te = new TextEditor()
                te.redo()

                expect(te.undoStack.length).toEqual(0)
                expect(te.redoStack.length).toEqual(0)
            })
        })
        describe("Redo after undo on write first character", () => {
            it("should redo the writing", () => {
                let te = new TextEditor()
                te.write("a")
                te.undo()
                te.redo()
                let output = te.read()

                expect(output).toEqual("a")
                
                expect(te.undoStack.length).toEqual(1)
                expect(te.redoStack.length).toEqual(0)
            })
        })
        describe("Redo after undo on write multiple characters", () => {
            it("should redo the writing", () => {
                let te = new TextEditor()
                te.write('w')
				te.write('o')
				te.write('r')
				te.write('l')
				te.write('d')
                te.undo()
                te.redo()
                let output = te.read()

                expect(output).toEqual("world")
                expect(te.undoStack.length).toEqual(5)
                expect(te.redoStack.length).toEqual(0)
            })
        })
        describe("Multiple redo after multiple undo on write multiple characters", () => {
            it("should redo the writing", () => {
                let te = new TextEditor()
                te.write('w')
				te.write('o')
				te.write('r')
				te.write('l')
				te.write('d')
                te.undo()
                te.undo()
                te.undo()
                te.redo()
                te.redo()
                let output = te.read()

                expect(output).toEqual("worl")
                expect(te.undoStack.length).toEqual(4)
                expect(te.redoStack.length).toEqual(1)
            })
        })
    })
})
