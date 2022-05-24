import useTodoStore from "../store/todoStore"
import { HStack, Text, Button } from "@chakra-ui/react"

export default function TodoItem({ id, text, isDone }) {
  // TODO: answer here

  return (
    <HStack justify="space-between" width="full">
      <Text aria-label="todo-text">
        {isDone && "✅ "}
        {text}
      </Text>
      <HStack>
        <Button
          size="xs"
          aria-label="delete-todo-button"
          backgroundColor="red.100"
          onClick={() => {
            removeTodo(id)
          }}
        >
          Delete
        </Button>
        <Button
          size="xs"
          aria-label={
            isDone ? "toggle-todo-undone-button" : "toggle-todo-done-button"
          }
          backgroundColor={isDone ? "yellow.300" : "green.300"}
          onClick={() => {
            toggleTodo(id)
          }}
        >
          {isDone ? "Undo" : "Done"}
        </Button>
      </HStack>
    </HStack>
  )
}
