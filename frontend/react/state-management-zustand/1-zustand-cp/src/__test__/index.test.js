import { render, screen, fireEvent } from "@testing-library/react"
import { act } from "react-dom/test-utils"
import App from "../App"

describe("Todo App test case", () => {
  it("should render the components", () => {
    render(<App />)
    const app = screen.getByLabelText("app")
    expect(app).toBeInTheDocument()
  })

  it("should add a new todo", () => {
    render(<App />)
    const addTodoInput = screen.getByLabelText("add-todo-input")
    const addTodoButton = screen.getByLabelText("add-todo-button")
    expect(addTodoInput).toBeInTheDocument()
    expect(addTodoButton).toBeInTheDocument()

    fireEvent.change(addTodoInput, { target: { value: "Hello World" } })
    act(() => {
      addTodoButton.click()
    })

    const newTodo = screen.getAllByText("Hello World")
    expect(newTodo[0]).toBeInTheDocument()
  })

  it('should count the number of total todos', () => {
      render(<App />)
      const count = screen.getByLabelText('todo-count-total')
      expect(count).toBeInTheDocument()
      expect(count).toHaveTextContent('Total : '+1)
  })

  it("should change the todo status", () => {
    render(<App />)

    const todoCheckbox = screen.getAllByLabelText("toggle-todo-done-button")[0]
    expect(todoCheckbox).toBeInTheDocument()

    act(() => {
      todoCheckbox.click()
    })

    const todoChecked = screen.getAllByLabelText("toggle-todo-undone-button")[0]
    expect(todoChecked).toBeInTheDocument()
  })

  it("should count the number of done todos", () => {
    render(<App />)
    const count = screen.getByLabelText("todo-count-total")
    expect(count).toBeInTheDocument()
    expect(count).toHaveTextContent("Total : " + 1)

    const doneCount = screen.getByLabelText("todo-count-done")
    expect(doneCount).toBeInTheDocument()
    expect(doneCount).toHaveTextContent("Done : " + 1)
  })

  it("should delete the todo", () => {
    render(<App />)

    const deleteTodoButton = screen.getAllByLabelText("delete-todo-button")[0]
    expect(deleteTodoButton).toBeInTheDocument()

    act(() => {
      deleteTodoButton.click()
    })

    const todoDeleted = screen.queryByText("todo-text")
    expect(todoDeleted).not.toBeInTheDocument()

    const count = screen.getByLabelText("todo-count-total")
    expect(count).toBeInTheDocument()
    expect(count).toHaveTextContent("Total : " + 0)
  })

  it("should add multiple todos", () => {
    render(<App />)

    const addTodoInput = screen.getByLabelText("add-todo-input")
    const addTodoButton = screen.getByLabelText("add-todo-button")
    expect(addTodoInput).toBeInTheDocument()
    expect(addTodoButton).toBeInTheDocument()

    fireEvent.change(addTodoInput, { target: { value: "Hello World" } })
    act(() => {
      addTodoButton.click()
    })

    fireEvent.change(addTodoInput, { target: { value: "Hello World 2" } })
    act(() => {
      addTodoButton.click()
    })

    fireEvent.change(addTodoInput, { target: { value: "Hello World 3" } })
    act(() => {
      addTodoButton.click()
    })

    const newTodo = screen.getAllByText("Hello World")
    expect(newTodo[0]).toBeInTheDocument()


    const count = screen.getByLabelText("todo-count-total")
    expect(count).toBeInTheDocument()
    expect(count).toHaveTextContent("Total : " + 3)

    const doneCount = screen.getByLabelText("todo-count-done")
    expect(doneCount).toBeInTheDocument()
    expect(doneCount).toHaveTextContent("Done : " + 0)

    const todoCheckbox = screen.getAllByLabelText("toggle-todo-done-button")[0]
    expect(todoCheckbox).toBeInTheDocument()

    act(() => {
        todoCheckbox.click()
        }
    )

    const todoChecked = screen.getAllByLabelText("toggle-todo-undone-button")[0]
    expect(todoChecked).toBeInTheDocument()

    expect(doneCount).toHaveTextContent("Done : " + 1)

    const deleteTodoButton = screen.getAllByLabelText("delete-todo-button")[0]
    expect(deleteTodoButton).toBeInTheDocument()

    act(() => {
        deleteTodoButton.click()
        }
    )

    expect(count).toHaveTextContent("Total : " + 2)
  })
})
