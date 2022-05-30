import { HStack, VStack, Text, Box, Image, Button } from "@chakra-ui/react"
import { useRef, useState } from "react"
import useChatStore from "../store/chatStore"
import useUserStore from "../store/userStore"
import ChatItem from "./ChatItem"
import InputChat from "./InputChat"

export default function LeftChat() {
  const { name, image } = useUserStore((state) => state.leftUser)
  const [showNotification, setShowNotification] = useState(false)
  const ownName = useUserStore((state) => state.rightUser.name)
  const messages = useChatStore((state) => state.messages)
  const bottomMessageRef = useRef(null)
  const unsubscribeNewMessage = useChatStore.subscribe((state) => {
    if (
      state.messages.length > 0 &&
      state.messages[state.messages.length - 1].sender !== ownName
    ) {
      onNewMessage()
    }
  })

  const onNewMessage = () => {
    setShowNotification(true)
    setTimeout(() => {
      // After 10 seconds set the show value to false
      setShowNotification(false)
    }, 10 * 1000)
  }

  const onClickNotification = () => {
    setShowNotification(false)
    bottomMessageRef.current.scrollIntoView({ behavior: "smooth" })
  }

  return (
    <Box width="md" rounded="xl" backgroundColor="gray.100">
      <HStack
        backgroundColor="blue.100"
        padding="0.5rem"
        roundedTop="xl"
        justify="space-between"
      >
        <HStack>
          <Image
            borderRadius="full"
            boxSize="32px"
            fit="cover"
            src={image}
            alt={name}
          />
          <Text fontSize="1rem" fontWeight="semibold">
            {name}
          </Text>
        </HStack>
        <Button
          hidden={!showNotification}
          backgroundColor="green.200"
          rounded="xl"
          size="xs"
          onClick={onClickNotification}
        >
          New Message
        </Button>
      </HStack>
      <VStack height="30rem" padding="1rem" overflowX="auto">
        {messages.map((message, index) => (
          <ChatItem
            key={index}
            name={name}
            image={image}
            message={message.message}
            isOwn={message.sender === ownName}
          />
        ))}
        <div ref={bottomMessageRef} />
      </VStack>
      <InputChat sender={ownName} />
    </Box>
  )
}
