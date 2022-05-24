import useTodoStore from "../store/todoStore"
import { useState } from "react"
import { HStack, Input, Button } from "@chakra-ui/react"

export default function AddTodo() {
  // TODO: answer here

  const [text, setText] = useState("")
  const handleSubmit = (e) => {
    e.preventDefault()
    const newTodo = {
      id: Date.now(),
      text,
      isDone: false,
    }
    addTodo(newTodo)
    setText("")
  }

  return (
    <HStack width="md">
      <Input aria-label='add-todo-input' placeholder='Add new todo' value={text} onChange={(e) => setText(e.target.value)} />
      <Button aria-label='add-todo-button' onClick={handleSubmit}>Add</Button>
    </HStack>
  )
}
