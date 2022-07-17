import { Flex, Text, Image } from "@chakra-ui/react"

export default function ChatItem({ name, image, message, isOwn }) {
  return (
    <Flex direction={isOwn ? "row-reverse" : "row"} width="full" gap="0.5rem">
      {!isOwn && (
        <Image src={image} alt={name} rounded="xl" boxSize="24px" fit="cover" />
      )}
      <Text
        fontSize="sm"
        fontWeight="semibold"
        textAlign={isOwn ? "right" : "left"}
        padding="0.25rem"
        rounded="xl"
        backgroundColor="white"
      >
        {message}
      </Text>
    </Flex>
  )
}
