import useChatStore from "../store/chatStore"
import { useState } from "react"
import { HStack, Input, Button } from "@chakra-ui/react"

export default function InputChat({ sender }) {
  const [message, setMessage] = useState("")
  const { addMessage } = useChatStore()
  const submit = () => {
    addMessage({ sender, message })
    setMessage("")
  }

  return (
    <HStack padding='1rem'>
      <Input
        rounded='xl'
        backgroundColor='white'
        placeholder="Send message"
        value={message}
        onChange={(e) => {
          setMessage(e.target.value)
        }}
      />
      <Button backgroundColor='blue.100' onClick={submit}>
        Send
      </Button>
    </HStack>
  )
}
