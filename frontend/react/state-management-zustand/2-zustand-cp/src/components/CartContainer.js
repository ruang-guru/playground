import useCartStore from "../store/cartStore"
import {
  Button,
  Text,
  HStack,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
  Image,
  VStack
} from "@chakra-ui/react"

export default function CartContainer() {
  const { items, removeItem, changeQuantity } = useCartStore()
  const { isOpen, onOpen, onClose } = useDisclosure()

  return (
    <>
      <Button aria-label='cart-open-button' onClick={onOpen}>{`Cart ${items.length}`}</Button>
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Cart</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
              <VStack>
                {items.map((item) => (
                    <CartItem
                        key={item.id}
                        item={item}
                        onRemove={removeItem}
                        onChange={changeQuantity}
                    />
                ))}
              </VStack>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  )
}

function CartItem({item, onChange, onRemove}) {
    return (
      <HStack justify="space-between" width="full">
        <HStack gap='0.5rem'>
          <Image aria-label='cart-item-image' src={item.image} alt={item.image} fit="cover" boxSize="3rem" />
          <Text aria-label='cart-item-name'>{item.name}</Text>
        </HStack>
        <HStack gap="0.5rem">
          <HStack>
            <Button
              aria-label='cart-item-decrement-button'
              size="xs"
              onClick={() => onChange(item.id, item.quantity - 1)}
            >
              -
            </Button>
            <Text aria-label='cart-item-quantity' minWidth='2rem' textAlign='center'>{item.quantity}</Text>
            <Button
              aria-label='cart-item-increment-button'
              size="xs"
              onClick={() => onChange(item.id, item.quantity + 1)}
            >
              +
            </Button>
          </HStack>
          <Button aria-label='cart-item-remove-button' size="xs" onClick={() => onRemove(item.id)}>
            Remove
          </Button>
        </HStack>
      </HStack>
    )
}
