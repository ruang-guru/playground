import useItemStore from "../store/itemStore"
import { useState } from "react"
import {
  Button,
  Text,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
  Input,
  VStack,
} from "@chakra-ui/react"

export default function NewItem() {
  const [id, setId] = useState("")
  const [name, setName] = useState("")
  const [image, setImage] = useState("")
  const [price, setPrice] = useState("")
  const [stock, setStock] = useState("")

  const addItem = useItemStore((state) => state.addItem)
  const { isOpen, onOpen, onClose } = useDisclosure()

  return (
    <>
      <Button aria-label="add-item-modal-open" onClick={onOpen}>
        Add Item
      </Button>
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Add Item</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <VStack align="left">
              <Text>ID</Text>
              <Input
                aria-label="add-item-id-input"
                value={id}
                onChange={(e) => setId(e.target.value)}
              />
              <Text>Name</Text>
              <Input
                aria-label="add-item-name-input"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
              <Text>Image URL</Text>
              <Input
                aria-label="add-item-image-input"
                value={image}
                onChange={(e) => setImage(e.target.value)}
              />
              <Text>Price</Text>
              <Input
                type="number"
                aria-label="add-item-price-input"
                value={price}
                onChange={(e) => setPrice(e.target.value)}
              />
              <Text>Stock</Text>
              <Input
                type="number"
                aria-label="add-item-stock-input"
                value={stock}
                onChange={(e) => setStock(e.target.value)}
              />
            </VStack>
          </ModalBody>
          <Button
            aria-label="add-item-button"
            onClick={() => {
              addItem({
                id,
                name,
                image,
                price,
                stock,
              })
              onClose()
              setId("")
              setName("")
              setImage("")
              setPrice("")
              setStock("")
            }}
          >
            Add Item
          </Button>
        </ModalContent>
      </Modal>
    </>
  )
}
