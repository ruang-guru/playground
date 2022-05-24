import { Box, Text, Input, Button, Textarea, Flex } from "@chakra-ui/react"
import useStore from "../store/Store"
import { useState } from "react"

export default function ProfileEdit() {
  const { name, address, bio, setName, setAddress, setBio } = useStore()
  const [nameValue, setNameValue] = useState(name)
  const [addressValue, setAddressValue] = useState(address)
  const [bioValue, setBioValue] = useState(bio)

  const handleSubmit = (e) => {
    e.preventDefault()
    setName(nameValue)
    setAddress(addressValue)
    setBio(bioValue)
  }

  return (
    <div>
      {`RootComponent --> ChildrenComponent-2`}
      <Box
        textAlign="left"
        padding="2rem"
        rounded="2xl"
        border="2px"
        minWidth="xl"
      >
        <Text>Name:</Text>
        <Input
          type="text"
          value={nameValue}
          onChange={(e) => setNameValue(e.target.value)}
        />
        <Text>Address:</Text>
        <Textarea
          type="text"
          value={addressValue}
          onChange={(e) => setAddressValue(e.target.value)}
        />
        <Text>Bio:</Text>
        <Textarea
          type="text"
          value={bioValue}
          onChange={(e) => setBioValue(e.target.value)}
        />
        <Flex justify="center" paddingTop="1rem">
          <Button onClick={handleSubmit} type="button">
            Save
          </Button>
        </Flex>
      </Box>
    </div>
  )
}
