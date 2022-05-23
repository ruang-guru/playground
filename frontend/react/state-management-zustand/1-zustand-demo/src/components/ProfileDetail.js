import { Box, Text } from "@chakra-ui/react"
import useStore from "../store/Store"

export default function ProfileDetail() {
  const { name, address, bio } = useStore()

  return (
    <div>
      {`RootComponent --> ChildrenComponent-1`}
      <Box padding="1rem" rounded="xl" border="2px" maxWidth="xl">
        <Text fontSize="2rem" fontWeight="bold" textColor="green.800">
          {name}
        </Text>
        <Text>{address}</Text>
        <Text textColor="blue.700" as="i">
          {bio}
        </Text>
      </Box>
    </div>
  )
}
