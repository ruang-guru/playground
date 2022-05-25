import {
  useDisclosure,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalBody,
  ModalCloseButton,
  Button,
  HStack,
  Text,
  VStack,
  ModalHeader,
  Input,
  Divider,
} from "@chakra-ui/react"
import useChatStore from "../store/chatStore"
import useUserStore from "../store/userStore"
import { useState } from "react"

export default function Menu() {
  const { isOpen, onOpen, onClose } = useDisclosure()
  return (
    <>
      <Button backgroundColor='blue.100' onClick={onOpen}>Menu</Button>
      <MenuDialog isOpen={isOpen} onClose={onClose} />
    </>
  )
}

function MenuDialog({ isOpen, onClose, ...props }) {
  const deleteAllChat = useChatStore((state) => state.clearMessages)
  const { leftUser, rightUser, setLeftUser, setRightUser } = useUserStore()
  const [newLeftUser, setNewLeftUser] = useState(leftUser)
  const [newRightUser, setNewRightUser] = useState(rightUser)
  return (
    <Modal isOpen={isOpen} onClose={onClose} {...props}>
      <ModalOverlay />
      <ModalContent>
        <ModalCloseButton />
        <ModalHeader>Edit Menu</ModalHeader>
        <ModalBody>
          <HStack justify="space-between">
            <Text>Delete all chat :</Text>
            <Button
              onClick={() => {
                deleteAllChat()
                onClose()
              }}
              backgroundColor="red.300"
              rounded="xl"
              size="xs"
            >
              Delete
            </Button>
          </HStack>
          <Divider paddingBlock="0.3rem" />
          <VStack align="start">
            <HStack width="full" justify="space-between" paddingTop="0.5rem">
              <Text fontWeight="semibold">Edit Left User</Text>
              <Button
                onClick={() => {
                  setLeftUser(newLeftUser)
                }}
                backgroundColor="green.300"
                rounded="xl"
                size="xs"
              >
                save
              </Button>
            </HStack>
            <Text>Name:</Text>
            <Input
              value={newLeftUser.name}
              onChange={(e) => {
                setNewLeftUser({ ...newLeftUser, name: e.target.value })
              }}
            />
            <Text>Image URL:</Text>
            <Input
              value={newLeftUser.image}
              onChange={(e) => {
                setNewLeftUser({ ...newLeftUser, image: e.target.value })
              }}
            />
            <Divider paddingBlock="0.3rem" />
            <HStack width="full" justify="space-between" paddingTop="0.5rem">
              <Text fontWeight="semibold">Edit Right User</Text>
              <Button
                onClick={() => {
                  setRightUser(newRightUser)
                }}
                backgroundColor="green.300"
                rounded="xl"
                size="xs"
              >
                save
              </Button>
            </HStack>
            <Text>Name:</Text>
            <Input
              value={newRightUser.name}
              onChange={(e) => {
                setNewRightUser({ ...newRightUser, name: e.target.value })
              }}
            />
            <Text>Image URL:</Text>
            <Input
              value={newRightUser.image}
              onChange={(e) => {
                setNewRightUser({ ...newRightUser, image: e.target.value })
              }}
            />
          </VStack>
        </ModalBody>
      </ModalContent>
    </Modal>
  )
}
