import './App.css';
import { HStack, Text } from '@chakra-ui/react';
import LeftChat from './components/LeftChatContainer';
import RightChat from './components/RightChatContainer';
import Menu from './components/MenuContainer';

function App() {
  return (
    <div className="App">
      <HStack
        justify="space-between"
        padding="0.5rem"
        width="full"
        backgroundColor="blue.400"
      >
        <Text fontSize="1.5rem" fontWeight="semibold">
          Jomblo Chat App
        </Text>
        <Menu />
      </HStack>
      <HStack align="start" width="full" justify="space-around" padding="1rem">
        <LeftChat />
        <RightChat />
      </HStack>
    </div>
  )
}

export default App;
